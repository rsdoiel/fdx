// Package cli holds the command line behavior of fdx2txt and txt2fdx so that
// plain go test covers it. The programs under cmd/ only call Fdx2txt and Txt2fdx.
package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	"github.com/rsdoiel/fdx"
)

// The exit codes of every command line tool in the workspace (DR-0003).
const (
	ExitOK         = 0
	ExitUsage      = 2
	ExitData       = 65
	ExitNoInput    = 66
	ExitInternal   = 70
	ExitCantCreate = 73
	ExitIO         = 74
	ExitNoPerm     = 77
)

// options are the options every program takes.
type options struct {
	help, license, version bool
	newline, quiet         bool
	input, output          string
}

// newFlagSet makes a FlagSet that reports its errors through the program and
// registers the standard options (newline is the default for -newline), each in a short and a long spelling.
func newFlagSet(appName string, o *options, newline bool) *flag.FlagSet {
	fset := flag.NewFlagSet(appName, flag.ContinueOnError)
	fset.SetOutput(io.Discard)
	fset.BoolVar(&o.help, "h", false, "display help")
	fset.BoolVar(&o.help, "help", false, "display help")
	fset.BoolVar(&o.license, "l", false, "display license")
	fset.BoolVar(&o.license, "license", false, "display license")
	fset.BoolVar(&o.version, "v", false, "display version")
	fset.BoolVar(&o.version, "version", false, "display version")
	fset.BoolVar(&o.newline, "nl", newline, "add a trailing newline")
	fset.BoolVar(&o.newline, "newline", newline, "add a trailing newline")
	fset.BoolVar(&o.quiet, "quiet", false, "suppress error messages")
	fset.StringVar(&o.input, "i", "", "input filename")
	fset.StringVar(&o.input, "input", "", "input filename")
	fset.StringVar(&o.output, "o", "", "output filename")
	fset.StringVar(&o.output, "output", "", "output filename")
	return fset
}

// fail writes one plain line, prefixed with the program name, to stderr unless
// the program was run with -quiet, and returns code (DR-0011).
func fail(appName string, o *options, stderr io.Writer, code int, format string, a ...any) int {
	if !o.quiet {
		fmt.Fprintf(stderr, "%s: %s\n", appName, fmt.Sprintf(format, a...))
	}
	return code
}

// classify turns an error from reading a named input into an exit code.
func classify(err error) int {
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return ExitNoInput
	case errors.Is(err, fs.ErrPermission):
		return ExitNoPerm
	default:
		return ExitIO
	}
}

// start parses args, answers -help, -license and -version, and works out the
// input and output names from the options or the first and second positional
// arguments. It returns done true with the exit code when the program is to
// stop there.
func start(appName, helpText string, fset *flag.FlagSet, o *options, args []string, stdout, stderr io.Writer) (code int, done bool) {
	if err := fset.Parse(args); err != nil {
		// The message is the flag package's, without a trailing newline or usage.
		return fail(appName, o, stderr, ExitUsage, "%v (try %s -help)", err, appName), true
	}
	switch {
	case o.help:
		fmt.Fprintln(stdout, fdx.FmtHelp(helpText, appName, fdx.Version, fdx.ReleaseDate, fdx.ReleaseHash))
		return ExitOK, true
	case o.license:
		fmt.Fprintln(stdout, strings.TrimSpace(fdx.LicenseText))
		return ExitOK, true
	case o.version:
		fmt.Fprintf(stdout, "%s %s %s\n", appName, fdx.Version, fdx.ReleaseHash)
		return ExitOK, true
	}
	rest := fset.Args()
	if len(rest) > 2 {
		return fail(appName, o, stderr, ExitUsage, "surplus argument %q (try %s -help)", rest[2], appName), true
	}
	if o.input == "" && len(rest) > 0 {
		o.input = rest[0]
	}
	if o.output == "" && len(rest) > 1 {
		o.output = rest[1]
	}
	return ExitOK, false
}

// readSource returns the bytes of the named file, or of stdin when name is
// empty or "-". On failure it returns the exit code.
func readSource(appName string, o *options, stdin io.Reader, stderr io.Writer) ([]byte, int) {
	if o.input != "" && o.input != "-" {
		src, err := os.ReadFile(o.input)
		if err != nil {
			return nil, fail(appName, o, stderr, classify(err), "%v", err)
		}
		return src, ExitOK
	}
	src, err := io.ReadAll(stdin)
	if err != nil {
		return nil, fail(appName, o, stderr, ExitIO, "%v", err)
	}
	return src, ExitOK
}

// write sends the result to the named output file, or to stdout when there is
// none or it is "-". The file is created only now, after the input has been
// read and converted, so a run that fails leaves an existing file alone. A
// trailing newline is added when the options ask for one.
func write(appName string, o *options, result []byte, stdout, stderr io.Writer) int {
	if o.newline {
		result = append(result, '\n')
	}
	if o.output == "" || o.output == "-" {
		if _, err := stdout.Write(result); err != nil {
			return fail(appName, o, stderr, ExitIO, "%v", err)
		}
		return ExitOK
	}
	f, err := os.Create(o.output)
	if err != nil {
		return fail(appName, o, stderr, ExitCantCreate, "%v", err)
	}
	if _, err := f.Write(result); err != nil {
		f.Close()
		return fail(appName, o, stderr, ExitIO, "%v", err)
	}
	if err := f.Close(); err != nil {
		return fail(appName, o, stderr, ExitIO, "%v", err)
	}
	return ExitOK
}
