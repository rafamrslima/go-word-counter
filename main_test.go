package main

import (
	"os"
	"testing"
)

func tempFile(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "test-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteString(content); err != nil {
		t.Fatal(err)
	}
	f.Close()
	return f.Name()
}

func TestCountWords(t *testing.T) {
	path := tempFile(t, "the quick brown fox jumps over the lazy dog")

	counts, total, err := countWords(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if total != 9 {
		t.Errorf("total: got %d, want 9", total)
	}
	if counts["the"] != 2 {
		t.Errorf("count of 'the': got %d, want 2", counts["the"])
	}
	if counts["fox"] != 1 {
		t.Errorf("count of 'fox': got %d, want 1", counts["fox"])
	}
}

func TestCountWords_CaseInsensitive(t *testing.T) {
	path := tempFile(t, "Go GO go gO")

	counts, total, err := countWords(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if total != 4 {
		t.Errorf("total: got %d, want 4", total)
	}
	if counts["go"] != 4 {
		t.Errorf("count of 'go': got %d, want 4", counts["go"])
	}
	if len(counts) != 1 {
		t.Errorf("unique words: got %d, want 1", len(counts))
	}
}

func TestCountWords_EmptyFile(t *testing.T) {
	path := tempFile(t, "")

	counts, total, err := countWords(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if total != 0 {
		t.Errorf("total: got %d, want 0", total)
	}
	if len(counts) != 0 {
		t.Errorf("unique words: got %d, want 0", len(counts))
	}
}

func TestCountWords_FileNotFound(t *testing.T) {
	_, _, err := countWords("nonexistent.txt")
	if err == nil {
		t.Error("expected error for missing file, got nil")
	}
}

func TestCountWords_MultipleLines(t *testing.T) {
	path := tempFile(t, "hello world\nhello go\nworld go go")

	counts, total, err := countWords(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if total != 7 {
		t.Errorf("total: got %d, want 7", total)
	}
	if counts["go"] != 3 {
		t.Errorf("count of 'go': got %d, want 3", counts["go"])
	}
	if counts["hello"] != 2 {
		t.Errorf("count of 'hello': got %d, want 2", counts["hello"])
	}
	if counts["world"] != 2 {
		t.Errorf("count of 'world': got %d, want 2", counts["world"])
	}
}
