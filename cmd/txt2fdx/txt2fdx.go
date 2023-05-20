//
// txt2fdx converts a plain text file into a fdx file.
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	// My packages
	"github.com/rsdoiel/fdx"
	"github.com/rsdoiel/fountain"
)

var (
	helpText = `%{app_name}(1) | {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name} 

# SYNOPSIS

{app_name} [OPTIONS]

# DESCRIPTION

{app_name} is a command line program that reads a plain text file file
and returns a fdx file.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-i
: read input for filename

-o
: write output to filename

-newline
: add a trailing newline

# EXAMPLES

Convert *screenplay.txt* into *screenplay.fdx*.

~~~
	{app_name} -i screenplay.txt -o screenplay.fdx
~~~

Or alternatively

~~~
    cat screenplay.txt | {app_name} > screenplay.fdx
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


func main() {
	appName := path.Base(os.Args[0])
	// NOTE: The following are set when version.go is generated
	version := fdx.Version
	releaseDate := fdx.ReleaseDate
	releaseHash := fdx.ReleaseHash
	fmtHelp := fdx.FmtHelp

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&newLine, "newline", true, "add a trailing newline")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.StringVar(&inputFName, "i", "", "set the input filename")
	flag.StringVar(&outputFName, "o", "", "set the output filename")

	// Parse environment and options
	flag.Parse()
	//args := flag.Args()

	// Setup IO
	var err error

	in := os.Stdin
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
		fmt.Fprintf(out, "%s %s %s\n", appName, version, releaseHash)
		os.Exit(0)
	}

	// ReadAll of input
	src, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}
	
	// Parse input
	screenplay, err := fountain.Parse(src)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}

	// Now create our fdx document
	document := fdx.NewFinalDraft()
	document.FromFountain(screenplay)
	src, err = document.ToXML()
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}

	//and then render as a string
	if newLine {
		fmt.Fprintf(out, "%s\n", src)
	} else {
		fmt.Fprintf(out, "%s", src)
	}
}
