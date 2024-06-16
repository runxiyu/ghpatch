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

Public domain, or CC0-1.0 as an alternative. A parent grant applies as
indicated in the following.

### WAIVER OF PATENT RIGHTS

Notwithstanding the provisions of section 4, Affirmer acknowledges 
that the design codes, utility, technical figures, charts, methods,
processes, and all information related to the patent/invention arising out
of the Work, the sources, source codes, objects, object codes and all
system documentation relating thereto shall not be confidential and shall
be non-proprietary for the public worldwide. To the greatest extent
permitted by, but not in contravention of, applicable law, Affirmer hereby
overtly, fully, permanently, irrevocably and unconditionally waives,
abandons, and surrenders all of Affirmer's patent rights and associated
claims and causes of action, whether now known or unknown (including
existing as well as future claims and causes of action) throughout the
world, whether arising under federal law, state law, common law, foreign
law, or otherwise, for any patent/invention included by Affirmer in the
Work that are licensable by the Affirmer, both currently owned or
controlled and acquired or controlled in the future (the "Waiver of Patent
Rights"). The Waiver of Patent Rights shall not be subject to 
revocation, rescission, cancellation, termination, or any other legal 
or equitable action to disrupt the quiet enjoyment of the Work by the 
public as contemplated by Affirmer's express Statement of Purpose.

Should any part of the Waiver of Patent Rights for any reason be judged
legally invalid or ineffective under applicable law, then the Waiver of
Patent Rights shall be preserved to the maximum extent permitted taking
into account Affirmer's express Statement of Purpose. In addition, to the
extent the Waiver of Patent Rights is so judged, Affirmer hereby grants to
each affected person a royalty-free, non transferable, non sublicensable,
non exclusive, irrevocable and unconditional license to Affirmer's patent
rights, where such license applies only to those patents/inventions
included by Affirmer in the Work that are licensable by Affirmer, both
currently owned or controlled and acquired or controlled in the future,
that are necessarily infringed by making, having made, using, selling,
offering for sale, importing or transferring the Work, or otherwise
practicing any method or procedure in any patent/invention included by
Affirmer in the Work, in whole or in part, alone or in combination with or
included in any other Work (i) in all territories worldwide, (ii) for the
maximum duration provided by applicable law or treaty (including future
time extensions), (iii) in any current or future medium and for any number
of copies, and (iv) for any purpose whatsoever, including without
limitation commercial, advertising or promotional purposes (the "Patent
License"). The Patent License shall be deemed effective as of the date
this licensing document was applied by Affirmer to the Work. Should any
part of the Patent License for any reason be judged legally invalid or
ineffective under applicable law, such partial invalidity or
ineffectiveness shall not invalidate the remainder of the Patent License,
and in such case Affirmer hereby affirms that they will not (i) exercise
any of their patent rights in the Work or (ii) assert any associated
claims and causes of action with respect to the Work, in either case
contrary to Affirmer's express Statement of Purpose.

(This waiver was adapted from Public Discussion Draft #2 of the Worldwide
Public Domain Dedication.)

## Mirrors

- <https://git.sr.ht/~runxiyu/ghpatch> is the "canonical" repository
- <https://git.runxiyu.org/runxiyu/current/ghpatch.git> also exists and could
  be potentially be used for automatic deployment in the future
- <https://github.com/runxiyu/ghpatch> exists; however, pull requests are
  currently hooked to the GitHub endpoint of
  [Hybrid](https://git.sr.ht/~runxiyu/hybrid) rather than `ghpatch` itself
