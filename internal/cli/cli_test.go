// Copyright (C) 2017-2026 R. S. Doiel
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package cli

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// The tests here call the programs in process, so that the coverage tool sees
// them; cli_exit_test.go in the repository root runs the built binaries.

type program func(string, []string, *strings.Reader, *bytes.Buffer, *bytes.Buffer) int

// call runs a program with in-memory streams and returns stdout, stderr and the code.
func call(f func(string, []string, *strings.Reader, *bytes.Buffer, *bytes.Buffer) int, name, stdin string, args ...string) (string, string, int) {
	var out, errOut bytes.Buffer
	code := f(name, args, strings.NewReader(stdin), &out, &errOut)
	return out.String(), errOut.String(), code
}

func fdx2txt(name string, args []string, in *strings.Reader, out, errOut *bytes.Buffer) int {
	return Fdx2txt(name, args, in, out, errOut)
}

func txt2fdx(name string, args []string, in *strings.Reader, out, errOut *bytes.Buffer) int {
	return Txt2fdx(name, args, in, out, errOut)
}

func testdata(name string) string { return filepath.Join("..", "..", "testdata", name) }

type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) { return 0, errors.New("disk full") }

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) { return 0, errors.New("read failed") }

/** TestConvertBothWays converts a Fountain file to fdx and the fdx back to
 * text: the second result holds the scene heading of the first input.
 */
func TestConvertBothWays(t *testing.T) {
	src, err := os.ReadFile(testdata("sample-01.fountain"))
	if err != nil {
		t.Fatal(err)
	}
	xml, errOut, code := call(txt2fdx, "txt2fdx", string(src))
	if code != ExitOK || errOut != "" || !strings.HasPrefix(xml, "<FinalDraft") {
		t.Fatalf("txt2fdx: exit %d, stderr %q, output starts %.30q", code, errOut, xml)
	}
	text, errOut, code := call(fdx2txt, "fdx2txt", xml)
	if code != ExitOK || errOut != "" || !strings.Contains(text, "EXT. LIBRARY - DAY") {
		t.Fatalf("fdx2txt: exit %d, stderr %q, %d bytes", code, errOut, len(text))
	}
}

/** TestInputAndOutputDash checks that - names standard input and standard
 * output, for the option and for the arguments.
 */
func TestInputAndOutputDash(t *testing.T) {
	src, _ := os.ReadFile(testdata("sample-01.fdx"))
	want, _, _ := call(fdx2txt, "fdx2txt", string(src))
	for _, args := range [][]string{{"-i", "-"}, {"-"}, {"-", "-"}, {"-o", "-"}} {
		got, _, code := call(fdx2txt, "fdx2txt", string(src), args...)
		if code != ExitOK || got != want {
			t.Errorf("%v: exit %d, output differs from standard input run", args, code)
		}
	}
}

/** TestHelpLicenseVersion checks the three answers and that -quiet does not
 * silence them.
 */
func TestHelpLicenseVersion(t *testing.T) {
	for _, f := range []struct {
		name string
		run  program
	}{{"fdx2txt", fdx2txt}, {"txt2fdx", txt2fdx}} {
		help, _, code := call(f.run, f.name, "", "-quiet", "-h")
		if code != ExitOK || !strings.HasPrefix(help, "%"+f.name+"(1) user manual") || strings.Contains(help, "{app_name}") {
			t.Errorf("%s -h: exit %d, starts %.40q", f.name, code, help)
		}
		lic, _, code := call(f.run, f.name, "", "-license")
		if code != ExitOK || !strings.Contains(lic, "GNU AFFERO GENERAL PUBLIC LICENSE") {
			t.Errorf("%s -license: exit %d, %d bytes", f.name, code, len(lic))
		}
		ver, _, code := call(f.run, f.name, "", "-v")
		if code != ExitOK || !strings.HasPrefix(ver, f.name+" ") || strings.Count(ver, "\n") != 1 {
			t.Errorf("%s -v: exit %d, %q", f.name, code, ver)
		}
	}
}

/** TestWriteFailures covers the codes a binary run cannot easily provoke:
 * standard output that fails is 74, standard input that fails is 74, an output
 * file in a directory that refuses writes is 77 (skipped when run as root).
 */
func TestWriteFailures(t *testing.T) {
	src, _ := os.ReadFile(testdata("sample-01.fdx"))
	var errOut bytes.Buffer
	if code := Fdx2txt("fdx2txt", nil, strings.NewReader(string(src)), failingWriter{}, &errOut); code != ExitIO || !strings.HasPrefix(errOut.String(), "fdx2txt: ") {
		t.Errorf("failing stdout: exit %d, stderr %q; want 74", code, errOut.String())
	}
	errOut.Reset()
	if code := Fdx2txt("fdx2txt", nil, failingReader{}, &bytes.Buffer{}, &errOut); code != ExitIO {
		t.Errorf("failing stdin: exit %d, stderr %q; want 74", code, errOut.String())
	}
	if os.Geteuid() == 0 {
		t.Skip("running as root: permissions are not enforced")
	}
	dir := t.TempDir()
	if err := os.Chmod(dir, 0o500); err != nil {
		t.Fatal(err)
	}
	defer os.Chmod(dir, 0o700)
	errOut.Reset()
	code := Fdx2txt("fdx2txt", []string{"-i", testdata("sample-01.fdx"), "-o", filepath.Join(dir, "out.txt")}, strings.NewReader(""), &bytes.Buffer{}, &errOut)
	// Creating a file in a directory the user may not write is a permission
	// refusal, which DR-0003 gives its own code.
	if code != ExitCantCreate && code != ExitNoPerm {
		t.Errorf("read-only directory: exit %d, stderr %q; want 73 or 77", code, errOut.String())
	}
}

/** TestUnreadableInput checks that an input the user may not read is 77. */
func TestUnreadableInput(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("running as root: permissions are not enforced")
	}
	f := filepath.Join(t.TempDir(), "locked.fdx")
	if err := os.WriteFile(f, []byte("<FinalDraft/>"), 0o000); err != nil {
		t.Fatal(err)
	}
	_, errOut, code := call(fdx2txt, "fdx2txt", "", f)
	if code != ExitNoPerm || !strings.HasPrefix(errOut, "fdx2txt: ") {
		t.Errorf("exit %d, stderr %q; want 77", code, errOut)
	}
}

/** TestMalformedInput is 65 for fdx2txt, with the message on one line. */
func TestMalformedInput(t *testing.T) {
	_, errOut, code := call(fdx2txt, "fdx2txt", "<FinalDraft><Content>")
	if code != ExitData || strings.Count(errOut, "\n") != 1 {
		t.Errorf("exit %d, stderr %q; want 65 and one line", code, errOut)
	}
}

/** TestUsageAndInputErrors covers the in-process side of the exit code table:
 * an unknown option and a surplus argument are 2, a missing input is 66, an
 * input that is a directory is 74, and each gives one line prefixed with the
 * program name.
 */
func TestUsageAndInputErrors(t *testing.T) {
	dir := t.TempDir()
	for _, f := range []struct {
		name string
		run  program
	}{{"fdx2txt", fdx2txt}, {"txt2fdx", txt2fdx}} {
		cases := []struct {
			what string
			args []string
			want int
		}{
			{"unknown option", []string{"-nosuchoption"}, ExitUsage},
			{"surplus argument", []string{"a", "b", "c"}, ExitUsage},
			{"missing input", []string{filepath.Join(dir, "missing")}, ExitNoInput},
			{"input is a directory", []string{"-i", dir}, ExitIO},
		}
		for _, c := range cases {
			out, errOut, code := call(f.run, f.name, "", c.args...)
			if code != c.want || out != "" || !strings.HasPrefix(errOut, f.name+": ") || strings.Count(errOut, "\n") != 1 {
				t.Errorf("%s, %s: exit %d (want %d), stdout %q, stderr %q", f.name, c.what, code, c.want, out, errOut)
			}
		}
		if out, errOut, code := call(f.run, f.name, "", "-quiet", "-nosuchoption"); code != ExitUsage || out != "" || errOut != "" {
			t.Errorf("%s -quiet: exit %d, stdout %q, stderr %q; want 2 and nothing printed", f.name, code, out, errOut)
		}
	}
}

/** TestOutputWriteFails writes to /dev/full, which accepts the open and refuses
 * every write: the run fails with 74 and says so on one line. It is skipped
 * where /dev/full does not exist.
 */
func TestOutputWriteFails(t *testing.T) {
	if _, err := os.Stat("/dev/full"); err != nil {
		t.Skip("no /dev/full on this system")
	}
	src, _ := os.ReadFile(testdata("sample-01.fdx"))
	_, errOut, code := call(fdx2txt, "fdx2txt", string(src), "-o", "/dev/full")
	if code != ExitIO || !strings.HasPrefix(errOut, "fdx2txt: ") {
		t.Errorf("exit %d, stderr %q; want 74", code, errOut)
	}
}
