
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
