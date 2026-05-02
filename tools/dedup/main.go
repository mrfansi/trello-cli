package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var typeDecl = regexp.MustCompile(`^type ([A-Za-z_][A-Za-z0-9_]*)[\s=]`)

// Anon-union path-param signatures generated from oneOf path params.
// Pattern across multiple lines:
//
//	<name> struct {
//	\t\tunion json.RawMessage
//	\t}
//
// Replace with `<name> string` so callers can pass plain IDs.
// Match anon-union only in function-param position: preceded by `(` or `, `.
var anonUnion = regexp.MustCompile(`([(,] ?)([A-Za-z_][A-Za-z0-9_]*) struct \{\s*union json\.RawMessage\s*\}`)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: dedup <file>")
		os.Exit(2)
	}
	path := os.Args[1]
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	deduped := dedupeTypes(data)
	rewritten := anonUnion.ReplaceAll(deduped, []byte("$1$2 string"))

	if err := os.WriteFile(path, rewritten, 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func dedupeTypes(in []byte) []byte {
	scanner := bufio.NewScanner(bytes.NewReader(in))
	scanner.Buffer(make([]byte, 1024*1024), 16*1024*1024)
	var out strings.Builder
	seen := map[string]bool{}
	skipping := false
	depth := 0

	for scanner.Scan() {
		line := scanner.Text()
		if skipping {
			depth += strings.Count(line, "{") - strings.Count(line, "}")
			if depth <= 0 {
				skipping = false
			}
			continue
		}
		if m := typeDecl.FindStringSubmatch(line); m != nil {
			name := m[1]
			if seen[name] {
				if strings.Contains(line, "{") && !strings.Contains(line, "}") {
					skipping = true
					depth = strings.Count(line, "{") - strings.Count(line, "}")
				}
				continue
			}
			seen[name] = true
		}
		out.WriteString(line)
		out.WriteByte('\n')
	}
	return []byte(out.String())
}
