package version

import (
	"strings"
	"sync"
	"testing"
)

func TestStringFormat(t *testing.T) {
	got := String()
	if !strings.Contains(got, " (") || !strings.HasSuffix(got, ")") {
		t.Errorf("String() %q does not match `<ver> (<commit>, <date>)`", got)
	}
}

func TestResolveFromLdflags(t *testing.T) {
	saveVersion, saveCommit, saveDate := Version, Commit, Date
	saveOnce := resolveOnce
	t.Cleanup(func() {
		Version, Commit, Date = saveVersion, saveCommit, saveDate
		resolveOnce = saveOnce
	})

	Version = "v9.9.9"
	Commit = "deadbeef"
	Date = "2099-01-01T00:00:00Z"
	resolveOnce = sync.Once{}

	got := String()
	want := "v9.9.9 (deadbeef, 2099-01-01T00:00:00Z)"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
