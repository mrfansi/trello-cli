// Package version exposes build-time metadata. Values are injected via
// -ldflags during `make build` / `make install` / goreleaser. When the
// binary is built without ldflags (e.g., `go install` against a module
// proxy), the package falls back to runtime/debug.ReadBuildInfo, which
// the Go toolchain stamps with module version + VCS revision + commit
// time on plain `go build` / `go install` since Go 1.18.
package version

import (
	"runtime/debug"
	"sync"
)

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

var resolveOnce sync.Once

// resolve fills in any sentinel values from runtime/debug.ReadBuildInfo
// the first time a version field is read.
func resolve() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	if Version == "dev" && info.Main.Version != "" && info.Main.Version != "(devel)" {
		Version = info.Main.Version
	}
	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			if Commit == "none" && s.Value != "" {
				if len(s.Value) > 7 {
					Commit = s.Value[:7]
				} else {
					Commit = s.Value
				}
			}
		case "vcs.time":
			if Date == "unknown" && s.Value != "" {
				Date = s.Value
			}
		case "vcs.modified":
			if Commit != "none" && s.Value == "true" {
				Commit += "-dirty"
			}
		}
	}
}

// String returns a single-line "vX.Y.Z (commit, date)" summary.
func String() string {
	resolveOnce.Do(resolve)
	return Version + " (" + Commit + ", " + Date + ")"
}
