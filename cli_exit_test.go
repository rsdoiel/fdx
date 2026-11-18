package fdx

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// These tests run the two programs as built binaries, the way a user does,
// with real files from testdata.

var programNames = []string{"fdx2txt", "txt2fdx"}

// programInput names the testdata file each program reads.
var programInput = map[string]string{
	"fdx2txt": "testdata/sample-01.fdx",
	"txt2fdx": "testdata/sample-01.fountain",
}

var builtDir string

// binaries builds the two programs once per test run.
func binaries(t *testing.T) string {
	t.Helper()
	if builtDir != "" {
		return builtDir
	}
	dir, err := os.MkdirTemp("", "fdx-cli-")
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range programNames {
		if out, err := exec.Command("go", "build", "-o", filepath.Join(dir, name), "./cmd/"+name).CombinedOutput(); err != nil {
			t.Fatalf("%s: build failed: %v\n%s", name, err, out)
		}
	}
	builtDir = dir
	return dir
}

func TestMain(m *testing.M) {
	code := m.Run()
	if builtDir != "" {
		os.RemoveAll(builtDir)
	}
	os.Exit(code)
}

// run runs a built program with stdin and returns stdout, stderr and the exit code.
func run(t *testing.T, dir, name, stdin string, args ...string) (string, string, int) {
	t.Helper()
	cmd := exec.Command(filepath.Join(binaries(t), name), args...)
	cmd.Dir = dir
	cmd.Stdin = strings.NewReader(stdin)
	var out, errOut bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &errOut
	code := 0
	if err := cmd.Run(); err != nil {
		var ee *exec.ExitError
		if !errors.As(err, &ee) {
			t.Fatalf("%s: %v", name, err)
		}
		code = ee.ExitCode()
	}
	return out.String(), errOut.String(), code
}

func readTestFile(t *testing.T, name string) string {
	t.Helper()
	b, err := os.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

/** TestOutputFileHoldsTheResult checks that -o writes the result to the file
 * and nothing to standard output, and that standard input, -i and the first
 * argument give the same result.
 */
func TestOutputFileHoldsTheResult(t *testing.T) {
	for _, name := range programNames {
		in, _ := filepath.Abs(programInput[name])
		want, _, code := run(t, t.TempDir(), name, "", "-i", in)
		if code != 0 || want == "" {
			t.Fatalf("%s -i: exit %d, %d bytes", name, code, len(want))
		}
		dir := t.TempDir()
		outFile := filepath.Join(dir, "out.txt")
		out, errOut, code := run(t, dir, name, "", "-i", in, "-o", outFile)
		if code != 0 || out != "" || readTestFile(t, outFile) != want {
			t.Errorf("%s -i -o: exit %d, stdout %d bytes, stderr %q; the file must hold the %d-byte result", name, code, len(out), errOut, len(want))
		}
		if out, _, _ := run(t, t.TempDir(), name, readTestFile(t, in)); out != want {
			t.Errorf("%s: standard input and -i differ", name)
		}
		if out, _, _ := run(t, t.TempDir(), name, "", in); out != want {
			t.Errorf("%s: positional input and -i differ", name)
		}
	}
}

/** TestNewlineDefaults pins the trailing newline: fdx2txt adds none unless
 * -newline is given, txt2fdx adds one unless -newline=false is given.
 */
func TestNewlineDefaults(t *testing.T) {
	for _, c := range []struct {
		name       string
		defaultNL  bool
		flagToggle string
	}{{"fdx2txt", false, "-newline"}, {"txt2fdx", true, "-newline=false"}} {
		in, _ := filepath.Abs(programInput[c.name])
		plain, _, _ := run(t, t.TempDir(), c.name, "", "-i", in)
		toggled, _, _ := run(t, t.TempDir(), c.name, "", "-i", in, c.flagToggle)
		with, without := plain, toggled
		if !c.defaultNL {
			with, without = toggled, plain
		}
		if with != without+"\n" {
			t.Errorf("%s: the newline option does not add exactly one newline (default has %d bytes, toggled %d)", c.name, len(plain), len(toggled))
		}
	}
}

/** TestStandardOptions checks -h, -l and -v in both spellings, before any file
 * is opened, and that the help page is the man page structure.
 */
func TestStandardOptions(t *testing.T) {
	for _, name := range programNames {
		for _, o := range []string{"-h", "-help", "--help"} {
			out, errOut, code := run(t, t.TempDir(), name, "", o, "-i", "no-such-file")
			if code != 0 || errOut != "" || !strings.HasPrefix(out, "%"+name+"(1) user manual | version ") || !strings.Contains(out, "# EXIT STATUS") {
				t.Errorf("%s %s: exit %d, stderr %q, help starts %.50q", name, o, code, errOut, out)
			}
		}
		for _, o := range []string{"-l", "-license", "-v", "-version"} {
			out, _, code := run(t, t.TempDir(), name, "", o)
			if code != 0 || out == "" {
				t.Errorf("%s %s: exit %d, %d bytes", name, o, code, len(out))
			}
		}
	}
}

/** TestExitCodes covers DR-0003 for every program: a missing input is 66, an
 * output that cannot be created is 73, a bad command line is 2, and a failing
 * run does not create or empty the output file.
 */
func TestExitCodes(t *testing.T) {
	for _, name := range programNames {
		dir := t.TempDir()
		missing := filepath.Join(dir, "missing.in")
		good, _ := filepath.Abs(programInput[name])
		noDir := filepath.Join(dir, "no", "such", "dir", "out")
		cases := []struct {
			what string
			args []string
			want int
		}{
			{"unknown option", []string{"-nosuchoption"}, 2},
			{"bad option value", []string{"-newline=maybe", good}, 2},
			{"surplus argument", []string{good, filepath.Join(dir, "o1"), filepath.Join(dir, "o2")}, 2},
			{"missing input file", []string{missing}, 66},
			{"missing input with -i", []string{"-i", missing}, 66},
			{"output cannot be created", []string{"-i", good, "-o", noDir}, 73},
		}
		for _, c := range cases {
			out, errOut, code := run(t, dir, name, "", c.args...)
			if code != c.want {
				t.Errorf("%s, %s: exit %d, want %d (stdout %.60q, stderr %q)", name, c.what, code, c.want, out, errOut)
			}
			if !strings.HasPrefix(errOut, name+": ") || strings.Count(errOut, "\n") != 1 {
				t.Errorf("%s, %s: stderr %q is not one line starting with the program name", name, c.what, errOut)
			}
		}
		outFile := filepath.Join(dir, "keep.txt")
		if err := os.WriteFile(outFile, []byte("keep"), 0o644); err != nil {
			t.Fatal(err)
		}
		run(t, dir, name, "", "-i", missing, "-o", outFile)
		if readTestFile(t, outFile) != "keep" {
			t.Errorf("%s: a missing input changed the existing output file", name)
		}
		out, errOut, code := run(t, dir, name, "", "-quiet", "-i", missing)
		if code != 66 || errOut != "" || out != "" {
			t.Errorf("%s -quiet: exit %d, stdout %q, stderr %q; want 66 and nothing printed", name, code, out, errOut)
		}
	}
}

/** TestMalformedFdx is exit 65: fdx2txt given a file that is not an fdx
 * document says so and leaves an existing output file alone.
 */
func TestMalformedFdx(t *testing.T) {
	dir := t.TempDir()
	outFile := filepath.Join(dir, "keep.txt")
	if err := os.WriteFile(outFile, []byte("keep"), 0o644); err != nil {
		t.Fatal(err)
	}
	out, errOut, code := run(t, dir, "fdx2txt", "<FinalDraft><Content>", "-o", outFile)
	if code != 65 || out != "" || !strings.HasPrefix(errOut, "fdx2txt: ") {
		t.Errorf("malformed fdx: exit %d, stdout %q, stderr %q; want 65", code, out, errOut)
	}
	if readTestFile(t, outFile) != "keep" {
		t.Error("a parse failure changed the existing output file")
	}
}
