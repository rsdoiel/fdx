
# Action Items

## Bugs

## Next

- [X] Remove dependency on caltechlibrary/cli
- [X] Update Makefile to generate version.go with Version, LicenseText and ReleaseDate (the date version.go was created)
- [X] Update Makefile so installer.sh and uninstaller.sh is created
- [X] Update codemeta-installer.sh to include Man pages
- [X] Generate individual man pages from -help
- [X] Update website build process to use website.mak
- [X] Add user-manual.md with links to documentation
- [X] Update INSTALL.md to include "Quick install with curl"
- [ ] Review XML output from ToXML() after a FromFountain() call, see where I need to add mapping for Text elements and embedded styling
- [ ] validate that I am producing fdx files that Final Draft, FadeIn and Trelby can read
- [ ] Add support to render as YAML

## Someday, Maybe

- [ ] Compile to WASM module then wrap for TypeScript
- [ ] Add fdx2json, json2fdx
- [ ] Add fdx2yaml, yaml2fdx
- [ ] Write and fdx2html using [scrippets](https://fountain.io/scrippets) approach
- [ ] Left/Right alignment should be respected based based on Paragraph Type
- [ ] Plaintext formatting needs to be pickup and respected from whole FinalDraft document (e.g. respect definitions, Layout, etc)
- [ ] Screen Headers and Footers can have Text, Dynamic, SceneProperties in any order, right now converting back to XML renders them in fixed order because they are ignored when rendering and plaintext

## Completed

- [x] Add ParseFile() to fdx.go
- [x] String() Paragraph needs to handle trailing new lines based on Paragraph Type
- [x] write tests that validate the source FDX content in _testout_ version

## Modernization (2026-11-16)

To match the workspace standard (knowledge and opml are the models; Go 1.27.3 across repositories). Do one repository at a time; run the tests and `make -n release` after each step.

- [x] Go 1.27.3: change `go 1.26.4` in go.mod, the `Go >=` lines in codemeta.json, README.md and INSTALL.md; run `go mod tidy`, `go vet ./...` and the tests.
- [x] Makefile: add `check-go-version` (go.mod, codemeta.json and the Go in use agree); `check-release` calls it.
- [x] Drop the armv7l target and 32-bit Pi wording (dropped 2026-10-27).
- [x] release.ps1: take knowledge's (existing-draft check, checksums).
- [x] Port DR-0012 (`check-committed`, `release` without `generate`, `release.bash --dry-run`, `makefile_test.go`).
- [x] CLI to `internal/cli` with `-h/-l/-v`, DR-0003 exit codes and Markdown help; `cli_exit_test.go`.
- [x] AGPL notice in every source header; `ioutil` removed; `/** */` blocks on every exported symbol; `CHANGES.md`.
- [x] fountain v1.0.4; version 1.0.5 prepared in `codemeta.json`.
- [ ] `DialogueBreaks` has a field named `XMLNAme` (typo for `XMLName`); renaming it is an API change, so decide whether it waits for a major version.

## Open questions that need Final Draft (or a bug report with an example file)

Found while writing tests (2026-11-18) by running `txt2fdx` on a script with every Fountain element type. Not tested as correct, because nothing here says what Final Draft does with them; the tests in `string_test.go` pin only what the code plainly intends.

- [ ] `FromFountain` writes `Paragraph.Type` from fountain's `TypeName()`: `Empty`, `Section`, `Synopsis`, `Page Feed`, `Boneyard` and `Lyric` are not Final Draft paragraph types as far as the `*Type` constants go (General, Scene Heading, Action, Character, Dialogue, Parenthetical, Transition, Shot, Cast List, Singing). Does Final Draft, Fade In or Trelby open a file with them? Should `Lyric` be `Singing`, and a `> centered <` line an Action with `Alignment="Center"` (it comes out as a Transition)? `Empty` paragraphs are written for every blank line.
- [ ] Title page: a Fountain key other than title, credit, author, copyright, draft date, contact, source and story by (for example `Notes:` or `Revision:`) is dropped silently, and every title page paragraph keeps the leading space of its Fountain value.
- [ ] `fountain` (not fdx): content before the first scene heading is dropped by `fountain.Parse`, so a script with only action lines gives zero elements. Check whether that is intended before relying on `txt2fdx` for fragments.
- [ ] `Paragraph.String` writes one line per `Text` run, so a paragraph whose words are split across runs (Final Draft does this wherever the style changes) comes out one run per line. Needs a real example file to see what Final Draft writes.

