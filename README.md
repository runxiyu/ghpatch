# ghpatch

A simple server that listens to pull requests from GitHub Webhooks and
converts them to emailed patches.

## Warning

```go
/*
NOTE: Do not use this in production yet as GitHub secrets validation,
      replay attack prevention, quite a lot of error checking, etc., are not
      implemented yet.
*/
```

## Configuration

Modify `main.go` to configure the program. (This will change once I'm a
bit more comfortable with Go.)

Run this behind a reverse proxy.

## Copyright status

Public domain, or CC0-1.0 as an alternative. A parent grant applies,
check the LICENSE file for details.
