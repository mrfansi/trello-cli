package cmdutil

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestDecodeSuccess(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"id":"abc","name":"Foo"}`)),
	}
	var got struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	if err := Decode(resp, &got); err != nil {
		t.Fatalf("Decode: %v", err)
	}
	if got.ID != "abc" || got.Name != "Foo" {
		t.Errorf("got %+v", got)
	}
}

func TestDecodeErrorStatus(t *testing.T) {
	resp := &http.Response{
		StatusCode: 401,
		Status:     "401 Unauthorized",
		Body:       io.NopCloser(strings.NewReader(`invalid token`)),
	}
	err := Decode(resp, nil)
	if err == nil {
		t.Fatal("want error, got nil")
	}
	if !strings.Contains(err.Error(), "401") || !strings.Contains(err.Error(), "invalid token") {
		t.Errorf("error missing context: %v", err)
	}
}

func TestDecodeNilTarget(t *testing.T) {
	resp := &http.Response{
		StatusCode: 204,
		Body:       io.NopCloser(strings.NewReader(``)),
	}
	if err := Decode(resp, nil); err != nil {
		t.Errorf("Decode(nil) on 2xx: %v", err)
	}
}

func TestContextHasTimeout(t *testing.T) {
	ctx, cancel := Context()
	defer cancel()
	deadline, ok := ctx.Deadline()
	if !ok {
		t.Fatal("context has no deadline")
	}
	if time.Until(deadline) <= 0 || time.Until(deadline) > 31*time.Second {
		t.Errorf("unexpected deadline: %v", deadline)
	}
	if ctx.Err() != nil {
		t.Errorf("context already errored: %v", ctx.Err())
	}
	_ = context.Background
}
