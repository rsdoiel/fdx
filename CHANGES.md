# Changes

Reconstructed for v1.0.3 and v1.0.4 from the tags and the git history; maintained going forward.

## v1.0.5 — 2026-11-18

### Changed

- **The licence is the GNU Affero General Public License, version 3 or (at your option) any later version, throughout.**
  `LICENSE` has said so since v1.0.3 (2024-05-20), but the headers of `fdx.go`, `fromfountain.go` and `fdx_test.go` still
  carried the BSD 2-Clause text of v1.0.2 and earlier, and no release note recorded the change. The headers now carry the AGPL
  notice, and `LicenseText` (shown by `-license`) is the AGPL text. v1.0.2 and the releases before it remain available under the
  BSD 2-Clause licence they were published with.
- **The two programs follow the workspace command line conventions.** Exit codes follow workspace DR-0003: 2 for a bad command
  line (an unknown option, a bad value, a surplus argument), 65 for input that is not an fdx file, 66 for a missing input,
  73 for an output that cannot be created, 74 for a read or write failure, 77 for a permission refusal. Before, every
  failure exited 1. An error is one line on standard error prefixed with the program name; `-quiet` drops the message and keeps
  the code. A script that tested for exit status 1 must test for non-zero.
- `-h`, `-l` and `-v` join `-help`, `-license` and `-version`; `-input`, `-output` and `-nl` are long and short spellings of
  `-i`, `-o` and `-newline`. The input and output may be given as the first and second argument; a third argument is an error
  (before, any argument was ignored and the program read standard input). `-help`, `-license` and `-version` answer before any
  file is opened. The trailing newline defaults are unchanged: `fdx2txt` adds none, `txt2fdx` adds one.
- A run that fails no longer creates or empties the output file; it is created after the input has been read and converted.
- Help text is Markdown in the man page structure with an EXIT STATUS section, kept in `internal/cli/helptext.go`; the command
  code moved to `internal/cli`. The `*.1.md` pages are generated from it by `make`.
- Every exported type, function and constant in the package has a documentation block with an example. The deprecated `io/ioutil`
  calls are gone. Dead commented-out code in `fdx.go` was removed.
- Tests: package coverage 54% to 99%, `internal/cli` 0% to 94% (run in process as well as through the built binaries). They pin what the
  code plainly intends (styles, paragraph rules, title page fields, nil handling, exit codes); they do not claim Final Draft agrees.
- Requires fountain v1.0.4. Go 1.27.3 is the minimum version. The Makefile gained `check-go-version` and `check-committed`; a
  release is built from a clean, pushed commit (workspace DR-0012); `release.bash` takes `--dry-run`; the armv7l target and the
  website's pagefind search are gone.

The package API is unchanged from v1.0.4.

## v1.0.4 — 2026-10-20

- Scene numbers survive a Fountain to FDX round trip. A Fountain heading ending `#12A#` becomes a Scene Heading with the FDX
  `Number` attribute set to `12A` and the text without the marker; a `Number` on a Scene Heading is written back to Fountain as
  `#12A#`. Requires fountain v1.0.3.

## v1.0.3 — 2024-05-20

- Changed the licence from BSD 2-Clause to AGPL-3.0-or-later (`LICENSE`).
