//
// fdx2txt converts a fdx file into plain text suitable to read from the console.
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	// My packages
	"github.com/rsdoiel/fdx"
)

var (
	helpText = `%{app_name}(1) | version {version} {release_hash}
% R. S. Doiel
% {release_date} 

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS]

# DESCRIPTION

{app_name} is a command line program that reads an fdx file
and returns plain text

# OPTIONS

-help
: display help

-license
: display license

-version
: display version


-i
: read input from file

-o
: write output to file

-newline
: add a trailing newline 

# EXAMPLES

Convert *screenplay.fdx* into *screenplay.txt*.

~~~
    {app_name} -i screenplay.fdx -o screenplay.txt
~~~

Or alternatively

~~~
    cat screenplay.fdx | fdx2txt > screenplay.txt
~~~

`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	newLine          bool
	quiet            bool
	inputFName       string
	outputFName      string
)


func fmtHelp(src string, appName string, version string, releaseDate string, releaseHash string) string {
	m := map[string]string{
		"{app_name}": appName,
		"{version}": version,
		"{release_date}": releaseDate,
		"{release_hash}": releaseHash,
	}
	for varname, val := range m {
		if strings.Contains(src, varname) {
			src = strings.ReplaceAll(src, varname, val)
		}
	}
	return src
}



func main() {
	appName := path.Base(os.Args[0])
	version := fdx.Version
	// NOTE: This is the date that version.go was generated.
	releaseDate := fdx.ReleaseDate
	releaseHash := fdx.ReleaseHash

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&newLine, "newline", false, "add a trailing newline")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.StringVar(&inputFName, "i", "", "set the input filename")
	flag.StringVar(&outputFName, "o", "", "set the output filename")

	// Parse environment and options
	flag.Parse()
	//args := flag.Args()

	// Setup IO
	var err error
	in  := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	if inputFName != "" {
		in, err = os.Open(inputFName)
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}
		defer in.Close()
	}
	if outputFName != "" {
		out, err = os.Create(outputFName)
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}
		defer out.Close()
	}

	// Process options
	if showHelp {
		fmt.Fprintf(out, "%s\n", fmtHelp(helpText, appName, version, releaseDate, releaseHash))
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(out, "%s\n", fdx.LicenseText)
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintln(out, "%s %s $s", appName, version, releaseHash)
		os.Exit(0)
	}

	// ReadAll of input
	src, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}

	// Parse input
	screenplay, err := fdx.Parse(src)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}

	//and then render as a string
	if newLine {
		fmt.Fprintf(out, "%s\n", screenplay.String())
	} else {
		fmt.Fprintf(out, "%s", screenplay.String())
	}
}
