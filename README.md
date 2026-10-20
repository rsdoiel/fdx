[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)

# fdx

_fdx_ is a Go package and two small command line programs for working with the XML files that [Final Draft](https://www.finaldraft.com/) writes (`.fdx`), and that other screenwriting programs such as [Trelby](http://www.trelby.org/), [Celtx](https://www.celtx.com/index.html) and [Fade In](https://www.fadeinpro.com/) read and write. The package parses an FDX document, writes it back as XML, renders it as Fountain-like plain text, and builds one from a [Fountain](https://fountain.io) document parsed by the [fountain](https://github.com/rsdoiel/fountain) package. Its types carry JSON and YAML tags.

## The package

```go
doc, err := fdx.ParseFile("screenplay.fdx")
fmt.Println(doc.String())              // Fountain-like plain text

screenplay, _ := fountain.ParseFile("screenplay.fountain")
out := fdx.NewFinalDraft()
out.FromFountain(screenplay)
xml, _ := out.ToXML()                  // an FDX document
```

- **Scene numbers.** A Fountain heading that ends in a number between hashes, `INT. KITCHEN - DAY #12A#`, becomes a Scene Heading whose `Number` attribute is "12A", and the heading text no longer holds the marker. Going the other way, a `Number` on a Scene Heading is written back as `#12A#`.
- **Boneyards.** A Fountain boneyard (`/* ... */`) becomes an ordinary paragraph, so it shows and prints in Final Draft.

## The programs

[fdx2txt](fdx2txt.1.md)
: reads an FDX file and writes plain text in a Fountain-like format.

[txt2fdx](txt2fdx.1.md)
: reads a Fountain file and writes an FDX file.

## Install

Get a release from [GitHub](https://github.com/rsdoiel/fdx/releases), or see [INSTALL.md](INSTALL.md), which also covers installing from source. To use the package in a Go program:

```shell
go get github.com/rsdoiel/fdx
```

## Build and test

You need Go 1.26.4 or later. Pandoc and GNU Make are only needed for the man pages and the website, and `cmt` (CMTools) only to regenerate `version.go`, `about.md` and `CITATION.cff` from `codemeta.json`.

```shell
make          # build the programs into bin/ and the man pages into man/
make test     # go test ./...
make check    # gofmt -l and go vet; rewrites nothing
```
