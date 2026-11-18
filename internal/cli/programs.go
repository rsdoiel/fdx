package cli

import (
	"io"

	"github.com/rsdoiel/fdx"
	"github.com/rsdoiel/fountain"
)

/** Fdx2txt is fdx2txt: it reads an fdx file and writes plain text.
 *
 * Parameters:
 *   appName (string) — the program name, for help and error messages
 *   args ([]string) — the command line without the program name
 *   stdin (io.Reader) — standard input
 *   stdout (io.Writer) — standard output
 *   stderr (io.Writer) — standard error
 *
 * Returns:
 *   int — the exit code
 *
 * Example:
 *   os.Exit(cli.Fdx2txt("fdx2txt", os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
 */
func Fdx2txt(appName string, args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	var o options
	fset := newFlagSet(appName, &o, false)
	if code, done := start(appName, Fdx2txtHelpText, fset, &o, args, stdout, stderr); done {
		return code
	}
	src, code := readSource(appName, &o, stdin, stderr)
	if code != ExitOK {
		return code
	}
	screenplay, err := fdx.Parse(src)
	if err != nil {
		return fail(appName, &o, stderr, ExitData, "%v", err)
	}
	return write(appName, &o, []byte(screenplay.String()), stdout, stderr)
}

/** Txt2fdx is txt2fdx: it reads a plain text (Fountain) screenplay and writes
 * an fdx file.
 *
 * Parameters:
 *   appName (string) — the program name, for help and error messages
 *   args ([]string) — the command line without the program name
 *   stdin (io.Reader) — standard input
 *   stdout (io.Writer) — standard output
 *   stderr (io.Writer) — standard error
 *
 * Returns:
 *   int — the exit code
 *
 * Example:
 *   os.Exit(cli.Txt2fdx("txt2fdx", os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
 */
func Txt2fdx(appName string, args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	var o options
	fset := newFlagSet(appName, &o, true)
	if code, done := start(appName, Txt2fdxHelpText, fset, &o, args, stdout, stderr); done {
		return code
	}
	src, code := readSource(appName, &o, stdin, stderr)
	if code != ExitOK {
		return code
	}
	screenplay, err := fountain.Parse(src)
	if err != nil {
		return fail(appName, &o, stderr, ExitData, "%v", err)
	}
	document := fdx.NewFinalDraft()
	document.FromFountain(screenplay)
	xml, err := document.ToXML()
	if err != nil {
		return fail(appName, &o, stderr, ExitInternal, "%v", err)
	}
	return write(appName, &o, xml, stdout, stderr)
}
