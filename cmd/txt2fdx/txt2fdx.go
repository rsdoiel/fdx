// txt2fdx reads a plain text screenplay and writes an fdx file. It is part of fdx; run txt2fdx -help.
package main

import (
	"os"
	"path"

	"github.com/rsdoiel/fdx/internal/cli"
)

func main() {
	os.Exit(cli.Txt2fdx(path.Base(os.Args[0]), os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
