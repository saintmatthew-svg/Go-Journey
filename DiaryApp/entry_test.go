package main

import (
	"testing"
)

func TestNewEntry(t *testing.T) {
	title := "Test Title"
	body := "Test Body"
	entry := NewEntry(1, title, body)

	if entry.ID != 1 {
		t.Errorf("Expected ID 1, got %d", entry.ID)
	}
	if entry.Title != title {
		t.Errorf("Expected title %s, got %s", title, entry.Title)
	}
	if entry.Body != body {
		t.Errorf("Expected body %s, got %s", body, entry.Body)
	}
	if entry.DateCreated.IsZero() {
		t.Error("DateCreated should not be zero")
	}
}

