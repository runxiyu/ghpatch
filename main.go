/*
This program is public domain, or under the terms of Creative Commons Zero
1.0 Universal, at your choice.  In addition, a Waiver of Patent Rights
apply.  See the LICENSE file for details.
*/

/*
NOTE: Do not use this in production yet as GitHub secrets validation,
      replay attack prevention, quite a lot of error checking, etc., are not
      implemented yet.
*/

/*
NOTE: This program requires a working git-send-email setup.  This should
      use /sbin/sendmail by default on unconfigured systems with sendmail.
      You could uncomment the lines that connect to a local SMTP server.
*/

package main

import (
	// "log"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

const ENVELOPE_FROM string = "hybrid@runxiyu.org"
const HEADER_FROM string = "Hybrid <hybrid@runxiyu.org>"
const REPLY_TO string = "me@runxiyu.org"

func getToAddress(repo string) string {
	switch repo {
	case "runxiyu/sjdb-src":
		return "~runxiyu/sjdb@lists.sr.ht"
	case "runxiyu/ykps-wsgi":
		return "~runxiyu/ykps@lists.sr.ht"
	default:
		return "~runxiyu/public-inbox@lists.sr.ht"
	}
}

type PullRequestMessage struct {
	Action     string
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
	PullRequest struct {
		PatchUrl string `json:"patch_url"`
	} `json:"pull_request"`
}

func pullRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// TODO: Authentication

	var jq PullRequestMessage
	err := json.NewDecoder(req.Body).Decode(&jq)
	if err != nil {
		panic(err)
	}

	repo := jq.Repository.FullName
	if len(repo) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No repo name found.")
		return
	}

	patchUrl := jq.PullRequest.PatchUrl
	if len(patchUrl) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No patch URL found.")
		return
	}

	f, err := os.CreateTemp("", "github-patch-*.mbox")
	defer os.Remove(f.Name())

	// TODO: make http request to get patch
	resp, err := http.Get(patchUrl)
	defer resp.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error fetching patch.")
		return
	}
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error fetching patch.")
		return
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error copying patch to temporary file.")
		return
	}

	err = exec.Command(
		"git", "send-email",
		"--from", HEADER_FROM,
		"--8bit-encoding", "UTF-8",
		"--to", getToAddress(repo),
		"--confirm", "never",
		// "--suppress-cc", "all",
		"--reply-to", REPLY_TO,
		"--envelope-sender", ENVELOPE_FROM,
		// "--no-smtp-auth",
		// "--smtp-server", "localhost",
		// "--smtp-server-port", "25",
		f.Name(),
	).Run()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error sending email.")
		return
	}
	if err := f.Close; f.Close() != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/github/pull-request", pullRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
