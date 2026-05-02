package main

import (
	"strings"
	"testing"
)

func TestDedupeTypes(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		notWant string
	}{
		{
			name: "drops second alias of same name",
			in: `package x
type Id1 = string
type Id1 = string
`,
			want:    "type Id1 = string\n",
			notWant: "type Id1 = string\ntype Id1 = string",
		},
		{
			name: "drops second struct of same name with body",
			in: `package x
type Foo struct {
	A int
}
type Foo struct {
	B int
}
`,
			want:    "type Foo struct {\n\tA int\n}",
			notWant: "B int",
		},
		{
			name: "keeps distinct types",
			in: `package x
type A = string
type B = int
`,
			want: "type A = string\ntype B = int\n",
		},
		{
			name: "ignores non-type lines",
			in: `package x
func F() {}
type A = string
`,
			want: "func F() {}\ntype A = string\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(dedupeTypes([]byte(tt.in)))
			if !strings.Contains(got, tt.want) {
				t.Errorf("missing %q in:\n%s", tt.want, got)
			}
			if tt.notWant != "" && strings.Contains(got, tt.notWant) {
				t.Errorf("unexpected %q in:\n%s", tt.notWant, got)
			}
		})
	}
}

func TestAnonUnionReplacement(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		notWant string
	}{
		{
			name: "replaces param after open paren",
			in: `func F(id struct {
	union json.RawMessage
}) {}`,
			want: "func F(id string)",
		},
		{
			name: "replaces param after comma",
			in: `func F(ctx Ctx, id struct {
	union json.RawMessage
}) {}`,
			want: "func F(ctx Ctx, id string)",
		},
		{
			name: "leaves type body declarations alone",
			in: `type Wrapper struct {
	id struct {
		union json.RawMessage
	}
}`,
			notWant: "id string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(process([]byte(tt.in)))
			if tt.want != "" && !strings.Contains(got, tt.want) {
				t.Errorf("missing %q in:\n%s", tt.want, got)
			}
			if tt.notWant != "" && strings.Contains(got, tt.notWant) {
				t.Errorf("unexpected %q in:\n%s", tt.notWant, got)
			}
		})
	}
}
