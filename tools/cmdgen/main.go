// cmdgen reads openapi.json and emits cobra subcommand files into
// internal/commands/auto/. One file per resource (first path segment),
// one cobra subcommand per operation. All operations are wired via the
// existing /raw helpers (path templating, query, body) — generated code
// stays small and uniform.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type spec struct {
	Paths map[string]pathItem `json:"paths"`
}

type pathItem struct {
	Parameters []parameter                `json:"parameters"`
	Operations map[string]json.RawMessage `json:"-"`
}

func (p *pathItem) UnmarshalJSON(data []byte) error {
	raw := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	p.Operations = map[string]json.RawMessage{}
	for k, v := range raw {
		if k == "parameters" {
			_ = json.Unmarshal(v, &p.Parameters)
			continue
		}
		p.Operations[k] = v
	}
	return nil
}

type operation struct {
	OperationID string      `json:"operationId"`
	Summary     string      `json:"summary"`
	Description string      `json:"description"`
	Parameters  []parameter `json:"parameters"`
	RequestBody *struct {
		Required bool `json:"required"`
	} `json:"requestBody"`
}

type parameter struct {
	Name     string `json:"name"`
	In       string `json:"in"`
	Required bool   `json:"required"`
	Schema   *struct {
		Type string `json:"type"`
	} `json:"schema"`
	Description string `json:"description"`
}

type genOp struct {
	Method      string
	Path        string
	OpID        string
	SubName     string
	GoFuncName  string
	PathParams  []parameter
	QueryParams []parameter
	HasBody     bool
	Summary     string
}

var httpMethods = map[string]bool{
	"get": true, "post": true, "put": true, "delete": true, "patch": true,
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: cmdgen <openapi.json> <code-out-dir> [--docs <docs-path>]")
		os.Exit(2)
	}
	specPath := os.Args[1]
	outDir := os.Args[2]
	docsPath := ""
	for i := 3; i < len(os.Args); i++ {
		if os.Args[i] == "--docs" && i+1 < len(os.Args) {
			docsPath = os.Args[i+1]
			i++
		}
	}

	groups, resources := loadGroups(specPath)

	if outDir != "-" {
		if err := os.MkdirAll(outDir, 0o755); err != nil {
			die("mkdir: %v", err)
		}
		for _, r := range resources {
			ops := groups[r]
			out := emitGroup(r, ops)
			path := filepath.Join(outDir, r+".go")
			if err := os.WriteFile(path, []byte(out), 0o644); err != nil {
				die("write %s: %v", path, err)
			}
		}
		root := emitRoot(resources)
		if err := os.WriteFile(filepath.Join(outDir, "auto.go"), []byte(root), 0o644); err != nil {
			die("write auto.go: %v", err)
		}
		fmt.Printf("generated %d resources, %d code files\n", len(resources), len(resources)+1)
	}

	if docsPath != "" {
		if err := os.MkdirAll(filepath.Dir(docsPath), 0o755); err != nil {
			die("mkdir docs: %v", err)
		}
		if err := os.WriteFile(docsPath, []byte(emitDocs(groups, resources)), 0o644); err != nil {
			die("write docs: %v", err)
		}
		fmt.Printf("wrote docs %s\n", docsPath)
	}
}

func loadGroups(specPath string) (map[string][]genOp, []string) {
	data, err := os.ReadFile(specPath)
	if err != nil {
		die("read spec: %v", err)
	}
	var s spec
	if err := json.Unmarshal(data, &s); err != nil {
		die("parse spec: %v", err)
	}

	groups := map[string][]genOp{}
	for path, item := range s.Paths {
		for method, raw := range item.Operations {
			if !httpMethods[method] {
				continue
			}
			var op operation
			if err := json.Unmarshal(raw, &op); err != nil {
				continue
			}
			op.Parameters = mergeParams(item.Parameters, op.Parameters)
			g := buildOp(strings.ToUpper(method), path, op)
			if g == nil {
				continue
			}
			res := resourceFromPath(path)
			groups[res] = append(groups[res], *g)
		}
	}

	resources := make([]string, 0, len(groups))
	for r := range groups {
		resources = append(resources, r)
	}
	sort.Strings(resources)

	for _, r := range resources {
		ops := groups[r]
		sort.Slice(ops, func(i, j int) bool { return ops[i].SubName < ops[j].SubName })
		dedupSubNames(ops)
		groups[r] = ops
	}
	return groups, resources
}

func die(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func resourceFromPath(p string) string {
	parts := strings.Split(strings.TrimPrefix(p, "/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		return "misc"
	}
	r := parts[0]
	r = sanitizeIdent(r)
	return r
}

var nonAlnum = regexp.MustCompile(`[^A-Za-z0-9]`)

func sanitizeIdent(s string) string {
	s = nonAlnum.ReplaceAllString(s, "")
	if s == "" {
		return "x"
	}
	return s
}

func mergeParams(pathLevel, opLevel []parameter) []parameter {
	seen := map[string]bool{}
	merged := make([]parameter, 0, len(pathLevel)+len(opLevel))
	for _, p := range opLevel {
		key := p.In + "/" + p.Name
		if seen[key] {
			continue
		}
		seen[key] = true
		merged = append(merged, p)
	}
	for _, p := range pathLevel {
		key := p.In + "/" + p.Name
		if seen[key] {
			continue
		}
		seen[key] = true
		merged = append(merged, p)
	}
	return merged
}

func buildOp(method, path string, op operation) *genOp {
	g := &genOp{
		Method:  method,
		Path:    path,
		OpID:    op.OperationID,
		Summary: op.Summary,
		HasBody: op.RequestBody != nil,
	}
	for _, p := range op.Parameters {
		switch p.In {
		case "path":
			g.PathParams = append(g.PathParams, p)
		case "query":
			g.QueryParams = append(g.QueryParams, p)
		}
	}
	g.SubName = subName(op.OperationID, method, path)
	g.GoFuncName = goFuncName(g.SubName)
	if g.SubName == "" || g.GoFuncName == "" {
		return nil
	}
	return g
}

func subName(opID, method, path string) string {
	if opID != "" {
		return strings.ToLower(opID)
	}
	parts := []string{strings.ToLower(method)}
	for _, seg := range strings.Split(strings.TrimPrefix(path, "/"), "/") {
		seg = strings.Trim(seg, "{}")
		seg = sanitizeIdent(seg)
		if seg != "" {
			parts = append(parts, strings.ToLower(seg))
		}
	}
	return strings.Join(parts, "-")
}

func goFuncName(subName string) string {
	parts := strings.Split(subName, "-")
	var b strings.Builder
	b.WriteString("cmd")
	for _, p := range parts {
		if p == "" {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]))
		if len(p) > 1 {
			b.WriteString(p[1:])
		}
	}
	return sanitizeIdent(b.String())
}

func dedupSubNames(ops []genOp) {
	seen := map[string]int{}
	for i := range ops {
		n := ops[i].SubName
		if c, ok := seen[n]; ok {
			seen[n] = c + 1
			ops[i].SubName = fmt.Sprintf("%s-%d", n, c+1)
			ops[i].GoFuncName = goFuncName(ops[i].SubName)
		} else {
			seen[n] = 1
		}
	}
}

func emitGroup(resource string, ops []genOp) string {
	var b strings.Builder
	b.WriteString("// Code generated by tools/cmdgen. DO NOT EDIT.\n\n")
	b.WriteString("package auto\n\n")
	b.WriteString("import \"github.com/spf13/cobra\"\n\n")

	groupFn := "Group" + strings.ToUpper(resource[:1]) + resource[1:]
	fmt.Fprintf(&b, "func %s() *cobra.Command {\n", groupFn)
	fmt.Fprintf(&b, "\tcmd := &cobra.Command{Use: %q, Short: %q}\n", resource, "Auto-generated commands for "+resource)
	for _, op := range ops {
		fmt.Fprintf(&b, "\tcmd.AddCommand(%s())\n", op.GoFuncName)
	}
	b.WriteString("\treturn cmd\n}\n\n")

	for _, op := range ops {
		emitOp(&b, op)
	}
	return b.String()
}

func emitOp(b *strings.Builder, op genOp) {
	short := op.Summary
	if short == "" {
		short = op.Method + " " + op.Path
	}
	short = sanitizeShort(short)

	fmt.Fprintf(b, "func %s() *cobra.Command {\n", op.GoFuncName)
	b.WriteString("\tvar (\n")
	for _, p := range op.QueryParams {
		fmt.Fprintf(b, "\t\tq_%s string\n", sanitizeIdent(p.Name))
	}
	if op.HasBody {
		b.WriteString("\t\tdataArg string\n")
	}
	b.WriteString("\t)\n")

	use := op.SubName
	for _, p := range op.PathParams {
		use += " <" + p.Name + ">"
	}

	fmt.Fprintf(b, "\tcmd := &cobra.Command{\n")
	fmt.Fprintf(b, "\t\tUse:   %q,\n", use)
	fmt.Fprintf(b, "\t\tShort: %q,\n", short)
	fmt.Fprintf(b, "\t\tArgs:  cobra.ExactArgs(%d),\n", len(op.PathParams))
	b.WriteString("\t\tRunE: func(cmd *cobra.Command, args []string) error {\n")
	fmt.Fprintf(b, "\t\t\tpathVars := []string{\n")
	for i, p := range op.PathParams {
		fmt.Fprintf(b, "\t\t\t\t%q + \"=\" + args[%d],\n", p.Name, i)
	}
	b.WriteString("\t\t\t}\n")

	b.WriteString("\t\t\tqueryVars := []string{}\n")
	for _, p := range op.QueryParams {
		id := sanitizeIdent(p.Name)
		fmt.Fprintf(b, "\t\t\tif q_%s != \"\" { queryVars = append(queryVars, %q + \"=\" + q_%s) }\n", id, p.Name, id)
	}

	dataExpr := `""`
	if op.HasBody {
		dataExpr = "dataArg"
	}
	fmt.Fprintf(b, "\t\t\treturn execRaw(cmd, %q, %q, pathVars, queryVars, %s)\n", op.Method, op.Path, dataExpr)
	b.WriteString("\t\t},\n\t}\n")

	for _, p := range op.QueryParams {
		id := sanitizeIdent(p.Name)
		desc := sanitizeShort(p.Description)
		fmt.Fprintf(b, "\tcmd.Flags().StringVar(&q_%s, %q, \"\", %q)\n", id, p.Name, desc)
	}
	if op.HasBody {
		b.WriteString("\tcmd.Flags().StringVar(&dataArg, \"data\", \"\", \"Request body: literal JSON or @file\")\n")
	}
	b.WriteString("\treturn cmd\n}\n\n")
}

func sanitizeShort(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", " ")
	if len(s) > 80 {
		s = s[:77] + "..."
	}
	return s
}

func emitDocs(groups map[string][]genOp, resources []string) string {
	var b strings.Builder
	b.WriteString("# trecli command reference\n\n")
	b.WriteString("Auto-generated from `openapi.json`. Do not edit by hand — re-run `make gen-cmds` to refresh.\n\n")

	totalOps := 0
	for _, r := range resources {
		totalOps += len(groups[r])
	}
	fmt.Fprintf(&b, "**Coverage**: %d resource groups, %d operations.\n\n", len(resources), totalOps)

	b.WriteString("## Resource groups\n\n")
	b.WriteString("| Group | Operations |\n|-------|-----------:|\n")
	for _, r := range resources {
		fmt.Fprintf(&b, "| [`%s`](#%s) | %d |\n", r, strings.ToLower(r), len(groups[r]))
	}
	b.WriteString("\n")

	b.WriteString("Plus two handcrafted commands:\n\n")
	b.WriteString("- `me` — alias for `members get-members-id me`.\n")
	b.WriteString("- `raw <METHOD> <PATH>` — passthrough to any endpoint.\n\n")

	for _, r := range resources {
		fmt.Fprintf(&b, "## %s\n\n", r)
		fmt.Fprintf(&b, "%d operations.\n\n", len(groups[r]))
		for _, op := range groups[r] {
			emitOpDoc(&b, r, op)
		}
	}
	return b.String()
}

func emitOpDoc(b *strings.Builder, resource string, op genOp) {
	fmt.Fprintf(b, "### `%s %s`\n\n", resource, op.SubName)
	fmt.Fprintf(b, "`%s %s`\n\n", op.Method, op.Path)
	if op.Summary != "" {
		fmt.Fprintf(b, "%s\n\n", strings.TrimSpace(op.Summary))
	}

	usage := "trecli " + resource + " " + op.SubName
	for _, p := range op.PathParams {
		usage += " <" + p.Name + ">"
	}
	fmt.Fprintf(b, "```bash\n%s\n```\n\n", usage)

	if len(op.PathParams) > 0 {
		b.WriteString("Path arguments:\n\n")
		for _, p := range op.PathParams {
			desc := strings.TrimSpace(p.Description)
			if desc == "" {
				desc = "(no description)"
			}
			fmt.Fprintf(b, "- `<%s>` — %s\n", p.Name, oneLine(desc))
		}
		b.WriteString("\n")
	}

	if len(op.QueryParams) > 0 {
		b.WriteString("Query flags:\n\n")
		for _, p := range op.QueryParams {
			desc := strings.TrimSpace(p.Description)
			if desc == "" {
				desc = "(no description)"
			}
			fmt.Fprintf(b, "- `--%s` — %s\n", p.Name, oneLine(desc))
		}
		b.WriteString("\n")
	}

	if op.HasBody {
		b.WriteString("Body: `--data <json|@file>` (optional JSON request body).\n\n")
	}
}

func oneLine(s string) string {
	s = strings.ReplaceAll(s, "\r\n", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	for strings.Contains(s, "  ") {
		s = strings.ReplaceAll(s, "  ", " ")
	}
	if len(s) > 200 {
		s = s[:197] + "..."
	}
	return s
}

func emitRoot(resources []string) string {
	var b strings.Builder
	b.WriteString("// Code generated by tools/cmdgen. DO NOT EDIT.\n\n")
	b.WriteString("package auto\n\n")
	b.WriteString("import \"github.com/spf13/cobra\"\n\n")
	b.WriteString("// Groups returns all auto-generated resource command groups.\n")
	b.WriteString("func Groups() []*cobra.Command {\n")
	b.WriteString("\treturn []*cobra.Command{\n")
	for _, r := range resources {
		fmt.Fprintf(&b, "\t\tGroup%s(),\n", strings.ToUpper(r[:1])+r[1:])
	}
	b.WriteString("\t}\n}\n")
	return b.String()
}
