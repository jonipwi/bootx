package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadRequestRejectsUnknownField(t *testing.T) {
	path := writeTempRequest(t, `{
  "request_id":"test",
  "capability_id":"assist.personal-decision.v1",
  "user_id":"owner-local",
  "created_at":"2026-07-12T10:00:00Z",
  "goal":"test",
  "question":"test?",
  "selected_content":"",
  "content_source":{"type":"test","origin_verified":false},
  "data_class":"personal",
  "declared_domain":"general",
  "memory_permission":"session",
  "remote_permission":"deny",
  "output_preference":"standard",
  "synthetic":false,
  "unexpected_authority":true
}`)
	if _, err := readRequest(path); err == nil || !strings.Contains(err.Error(), "unknown field") {
		t.Fatalf("expected unknown field error, got %v", err)
	}
}

func TestReadRequestRejectsMultipleObjects(t *testing.T) {
	path := writeTempRequest(t, `{}`+"\n"+`{}`)
	if _, err := readRequest(path); err == nil || !strings.Contains(err.Error(), "exactly one") {
		t.Fatalf("expected one-object error, got %v", err)
	}
}

func writeTempRequest(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "request.json")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}
