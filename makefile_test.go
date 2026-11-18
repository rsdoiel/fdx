package fdx

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

/** TestMakefileProgramsMatchManPages checks that the Makefile's PROGRAMS line
 * names every cmd/ directory and every section 1 man page source. The man page
 * list is derived from PROGRAMS, so a program left out would drop its page from
 * `make man` and `make install` while the build still succeeds.
 */
func TestMakefileProgramsMatchManPages(t *testing.T) {
	mk, err := os.ReadFile("Makefile")
	if err != nil {
		t.Fatal(err)
	}
	m := regexp.MustCompile(`(?m)^PROGRAMS *= *(.*)$`).FindSubmatch(mk)
	if m == nil {
		t.Fatal("Makefile has no PROGRAMS line")
	}
	got := strings.Fields(string(m[1]))
	var cmds, pages []string
	dirs, _ := filepath.Glob("cmd/*")
	for _, d := range dirs {
		cmds = append(cmds, filepath.Base(d))
	}
	md, _ := filepath.Glob("*.1.md")
	for _, f := range md {
		pages = append(pages, strings.TrimSuffix(f, ".1.md"))
	}
	sort.Strings(got)
	sort.Strings(cmds)
	sort.Strings(pages)
	if strings.Join(got, " ") != strings.Join(cmds, " ") {
		t.Errorf("Makefile PROGRAMS = %v, cmd/ holds %v", got, cmds)
	}
	if strings.Join(got, " ") != strings.Join(pages, " ") {
		t.Errorf("Makefile PROGRAMS = %v, *.1.md pages = %v", got, pages)
	}
}

/** TestCheckCommitted runs `make check-committed` in a scratch repository with
 * a bare origin: it passes when the tree is clean and HEAD is pushed, and fails
 * for an uncommitted change to a tracked file and for an unpushed commit.
 */
func TestCheckCommitted(t *testing.T) {
	for _, tool := range []string{"git", "make"} {
		if _, err := exec.LookPath(tool); err != nil {
			t.Skipf("%s not available", tool)
		}
	}
	mk, err := filepath.Abs("Makefile")
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	origin, work := filepath.Join(root, "origin.git"), filepath.Join(root, "work")
	run := func(dir string, name string, args ...string) error {
		c := exec.Command(name, args...)
		c.Dir = dir
		c.Env = append(os.Environ(), "GIT_AUTHOR_NAME=t", "GIT_AUTHOR_EMAIL=t@example.org",
			"GIT_COMMITTER_NAME=t", "GIT_COMMITTER_EMAIL=t@example.org")
		if b, err := c.CombinedOutput(); err != nil {
			t.Logf("%s %v: %s", name, args, b)
			return err
		}
		return nil
	}
	must := func(dir, name string, args ...string) {
		t.Helper()
		if err := run(dir, name, args...); err != nil {
			t.Fatalf("%s %v: %v", name, args, err)
		}
	}
	must(root, "git", "init", "-q", "--bare", origin)
	must(root, "git", "init", "-q", "-b", "main", work)
	must(work, "git", "remote", "add", "origin", origin)
	if err := os.WriteFile(filepath.Join(work, "a.txt"), []byte("one\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	must(work, "git", "add", "a.txt")
	must(work, "git", "commit", "-q", "-m", "one")
	must(work, "git", "push", "-q", "origin", "main")
	check := func() error { return run(work, "make", "-s", "-f", mk, "check-committed") }
	if err := check(); err != nil {
		t.Errorf("clean and pushed: want pass, got %v", err)
	}
	if err := os.WriteFile(filepath.Join(work, "a.txt"), []byte("two\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := check(); err == nil {
		t.Error("uncommitted change to a tracked file: want failure")
	}
	must(work, "git", "commit", "-q", "-am", "two")
	if err := check(); err == nil {
		t.Error("committed but unpushed: want failure")
	}
	must(work, "git", "push", "-q", "origin", "main")
	if err := check(); err != nil {
		t.Errorf("pushed again: want pass, got %v", err)
	}
}
