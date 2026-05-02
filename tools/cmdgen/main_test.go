package main

import (
	"reflect"
	"testing"
)

func TestResourceFromPath(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"/boards/{id}", "boards"},
		{"/boards/{id}/labels", "boards"},
		{"/customFields/{id}", "customFields"},
		{"/", "misc"},
		{"", "misc"},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			if got := resourceFromPath(tt.path); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSubName(t *testing.T) {
	tests := []struct {
		name   string
		opID   string
		method string
		path   string
		want   string
	}{
		{"uses operationId", "get-boards-id", "GET", "/boards/{id}", "get-boards-id"},
		{"lowercases operationId", "GET-Boards-Id", "GET", "/boards/{id}", "get-boards-id"},
		{"falls back to method+path", "", "GET", "/boards/{id}/labels", "get-boards-id-labels"},
		{"sanitizes path segments", "", "GET", "/foo-bar/{x}", "get-foobar-x"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subName(tt.opID, tt.method, tt.path); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestGoFuncName(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"get-boards-id", "cmdGetBoardsId"},
		{"post-cards", "cmdPostCards"},
		{"a-b-c", "cmdABC"},
		{"single", "cmdSingle"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := goFuncName(tt.in); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSanitizeIdent(t *testing.T) {
	tests := map[string]string{
		"foo":         "foo",
		"foo-bar":     "foobar",
		"customField": "customField",
		"123":         "123",
		"":            "x",
		"!@#":         "x",
	}
	for in, want := range tests {
		t.Run(in, func(t *testing.T) {
			if got := sanitizeIdent(in); got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestMergeParams(t *testing.T) {
	pathLevel := []parameter{
		{Name: "id", In: "path", Required: true},
	}
	opLevel := []parameter{
		{Name: "fields", In: "query"},
	}
	got := mergeParams(pathLevel, opLevel)
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	want := []parameter{
		{Name: "fields", In: "query"},
		{Name: "id", In: "path", Required: true},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestMergeParamsOpOverridesPath(t *testing.T) {
	pathLevel := []parameter{{Name: "id", In: "path", Description: "from path"}}
	opLevel := []parameter{{Name: "id", In: "path", Description: "from op"}}
	got := mergeParams(pathLevel, opLevel)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
	if got[0].Description != "from op" {
		t.Errorf("op-level should win, got %q", got[0].Description)
	}
}

func TestDedupSubNames(t *testing.T) {
	ops := []genOp{
		{SubName: "get", GoFuncName: "cmdGet"},
		{SubName: "get", GoFuncName: "cmdGet"},
		{SubName: "post", GoFuncName: "cmdPost"},
		{SubName: "get", GoFuncName: "cmdGet"},
	}
	dedupSubNames(ops)
	names := []string{ops[0].SubName, ops[1].SubName, ops[2].SubName, ops[3].SubName}
	want := []string{"get", "get-2", "post", "get-3"}
	if !reflect.DeepEqual(names, want) {
		t.Errorf("got %v, want %v", names, want)
	}
}
