package output

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func newTestRenderer(jsonMode bool) (*Renderer, *bytes.Buffer) {
	buf := &bytes.Buffer{}
	return &Renderer{JSON: jsonMode, W: buf}, buf
}

func TestTableRendersHeadersAndRows(t *testing.T) {
	r, buf := newTestRenderer(false)
	if err := r.Table([]string{"ID", "Name"}, [][]string{{"1", "Foo"}, {"2", "Bar"}}); err != nil {
		t.Fatalf("Table: %v", err)
	}
	out := buf.String()
	for _, want := range []string{"ID", "Name", "1", "Foo", "2", "Bar"} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in:\n%s", want, out)
		}
	}
}

func TestTableJSONMode(t *testing.T) {
	r, buf := newTestRenderer(true)
	if err := r.Table([]string{"ID", "Name"}, [][]string{{"1", "Foo"}}); err != nil {
		t.Fatalf("Table: %v", err)
	}
	var got []map[string]string
	if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal: %v\noutput: %s", err, buf.String())
	}
	if len(got) != 1 || got[0]["ID"] != "1" || got[0]["Name"] != "Foo" {
		t.Errorf("got %v", got)
	}
}

func TestKeyValueJSONMode(t *testing.T) {
	r, buf := newTestRenderer(true)
	if err := r.KeyValue([][2]string{{"ID", "1"}, {"Name", "Foo"}}); err != nil {
		t.Fatalf("KeyValue: %v", err)
	}
	var got map[string]string
	if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got["ID"] != "1" || got["Name"] != "Foo" {
		t.Errorf("got %v", got)
	}
}

func TestKeyValueAligns(t *testing.T) {
	r, buf := newTestRenderer(false)
	if err := r.KeyValue([][2]string{{"ID", "1"}, {"FullName", "Foo"}}); err != nil {
		t.Fatalf("KeyValue: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "ID") || !strings.Contains(out, "FullName") {
		t.Errorf("missing keys: %s", out)
	}
	if !strings.Contains(out, "Foo") || !strings.Contains(out, "1") {
		t.Errorf("missing values: %s", out)
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		name string
		in   string
		n    int
		want string
	}{
		{"short stays", "abc", 5, "abc"},
		{"newline replaced", "a\nb", 5, "a b"},
		{"truncated with ellipsis", "abcdef", 4, "abc…"},
		{"n=1 returns single char", "abc", 1, "a"},
		{"exact length stays", "abc", 3, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Truncate(tt.in, tt.n); got != tt.want {
				t.Errorf("Truncate(%q, %d) = %q, want %q", tt.in, tt.n, got, tt.want)
			}
		})
	}
}

func TestDeref(t *testing.T) {
	s := "hello"
	if got := Deref(&s); got != "hello" {
		t.Errorf("Deref(&s) = %q", got)
	}
	var nilPtr *string
	if got := Deref(nilPtr); got != "" {
		t.Errorf("Deref(nil) = %q, want empty", got)
	}
	n := 42
	if got := Deref(&n); got != 42 {
		t.Errorf("Deref(&42) = %d", got)
	}
	var nilInt *int
	if got := Deref(nilInt); got != 0 {
		t.Errorf("Deref(nil int) = %d, want 0", got)
	}
}
