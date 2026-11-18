// fdx2txt reads an fdx file and writes plain text. It is part of fdx; run fdx2txt -help.
package main

import (
	"os"
	"path"

	"github.com/rsdoiel/fdx/internal/cli"
)

func main() {
	os.Exit(cli.Fdx2txt(path.Base(os.Args[0]), os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
