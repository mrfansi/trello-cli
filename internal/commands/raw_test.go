package commands

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseMethodPath(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		methodFlag string
		wantMethod string
		wantPath   string
		wantErr    bool
	}{
		{"two args", []string{"GET", "/me"}, "", "GET", "/me", false},
		{"lowercase method upper", []string{"get", "/me"}, "", "GET", "/me", false},
		{"flag method", []string{"/me"}, "post", "POST", "/me", false},
		{"missing method", []string{"/me"}, "", "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, p, err := parseMethodPath(tt.args, tt.methodFlag)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err=%v wantErr=%v", err, tt.wantErr)
			}
			if err != nil {
				return
			}
			if m != tt.wantMethod || p != tt.wantPath {
				t.Errorf("got (%q, %q), want (%q, %q)", m, p, tt.wantMethod, tt.wantPath)
			}
		})
	}
}

func TestApplyPathVars(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		vars    []string
		want    string
		wantErr bool
	}{
		{"single sub", "/boards/{id}", []string{"id=abc"}, "/boards/abc", false},
		{"two subs", "/cards/{id}/checklists/{cid}", []string{"id=a", "cid=b"}, "/cards/a/checklists/b", false},
		{"escapes value", "/boards/{id}", []string{"id=a/b"}, "/boards/a%2Fb", false},
		{"unknown placeholder", "/boards/{id}", []string{"foo=x"}, "", true},
		{"unresolved placeholder", "/boards/{id}", nil, "", true},
		{"no placeholders", "/me", nil, "/me", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := applyPathVars(tt.path, tt.vars)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err=%v wantErr=%v", err, tt.wantErr)
			}
			if err != nil {
				return
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSplitKV(t *testing.T) {
	tests := []struct {
		in        string
		wantK     string
		wantV     string
		wantErr   bool
	}{
		{"k=v", "k", "v", false},
		{"key=val=with=equals", "key", "val=with=equals", false},
		{"empty=ok", "empty", "ok", false},
		{"=novalue", "", "", true},
		{"novalue", "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			k, v, err := splitKV(tt.in, "--test")
			if (err != nil) != tt.wantErr {
				t.Fatalf("err=%v wantErr=%v", err, tt.wantErr)
			}
			if err != nil {
				return
			}
			if k != tt.wantK || v != tt.wantV {
				t.Errorf("got (%q, %q), want (%q, %q)", k, v, tt.wantK, tt.wantV)
			}
		})
	}
}

func TestLoadBody(t *testing.T) {
	t.Run("empty returns nil", func(t *testing.T) {
		got, err := loadBody("")
		if err != nil {
			t.Fatal(err)
		}
		if got != nil {
			t.Errorf("got %q, want nil", got)
		}
	})
	t.Run("literal json passthrough", func(t *testing.T) {
		got, err := loadBody(`{"a":1}`)
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != `{"a":1}` {
			t.Errorf("got %q", got)
		}
	})
	t.Run("@file reads file", func(t *testing.T) {
		dir := t.TempDir()
		file := filepath.Join(dir, "body.json")
		if err := os.WriteFile(file, []byte(`{"hi":true}`), 0o600); err != nil {
			t.Fatal(err)
		}
		got, err := loadBody("@" + file)
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != `{"hi":true}` {
			t.Errorf("got %q", got)
		}
	})
	t.Run("@missing file errors", func(t *testing.T) {
		_, err := loadBody("@/nonexistent/file/xyz.json")
		if err == nil {
			t.Fatal("want error")
		}
	})
}

func TestBuildRequestQueryAndBody(t *testing.T) {
	body := []byte(`{"x":1}`)
	req, err := buildRequest("https://api.trello.com/1/cards", "POST", body,
		[]string{"idList=abc", "name=Foo"},
		[]string{"X-Trace=t1"},
	)
	if err != nil {
		t.Fatal(err)
	}
	if req.Method != "POST" {
		t.Errorf("method = %q", req.Method)
	}
	q := req.URL.Query()
	if q.Get("idList") != "abc" || q.Get("name") != "Foo" {
		t.Errorf("query missing: %v", q)
	}
	if req.Header.Get("Content-Type") != "application/json" {
		t.Errorf("content-type = %q", req.Header.Get("Content-Type"))
	}
	if req.Header.Get("X-Trace") != "t1" {
		t.Errorf("custom header missing")
	}
	got, _ := io.ReadAll(req.Body)
	if string(got) != `{"x":1}` {
		t.Errorf("body = %q", got)
	}
}

func TestBuildRequestNoBodyNoContentType(t *testing.T) {
	req, err := buildRequest("https://api.trello.com/1/me", "GET", nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if req.Header.Get("Content-Type") != "" {
		t.Errorf("unexpected content-type for empty body")
	}
}

func TestBuildRequestRejectsBadKV(t *testing.T) {
	_, err := buildRequest("https://api.trello.com/1/me", "GET", nil, []string{"badpair"}, nil)
	if err == nil || !strings.Contains(err.Error(), "--query") {
		t.Errorf("want --query error, got %v", err)
	}
}
