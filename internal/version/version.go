// Package version exposes build-time metadata. Values are injected via
// -ldflags during `make build` / `make install` / goreleaser. When the
// binary is built without ldflags (e.g., a plain `go build`), the
// fallback values below are used.
package version

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

// String returns a single-line "vX.Y.Z (commit, date)" summary.
func String() string {
	return Version + " (" + Commit + ", " + Date + ")"
}
