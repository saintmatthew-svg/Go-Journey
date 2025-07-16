package main

import (
	"testing"
)

func TestDiariesAdd(t *testing.T) {
	diaries := &Diaries{}
	diary := &Diary{username: "testuser", password: "testpass"}

	diaries.Add(diary)

	if len(diaries.diaries) != 1 {
		t.Errorf("Expected 1 diary, got %d", len(diaries.diaries))
	}

	if diaries.diaries[0] != diary {
		t.Errorf("Expected diary to be added, but it's not the same instance")
	}
}




func TestDiariesFindByUsername(t *testing.T) {
	diaries := &Diaries{}
	diary1 := &Diary{username: "user1", password: "pass1"}
	diary2 := &Diary{username: "user2", password: "pass2"}

	diaries.Add(diary1)
	diaries.Add(diary2)

	foundDiary := diaries.FindByUsername("user1")
	if foundDiary == nil || foundDiary.username != "user1" {
		t.Errorf("Expected to find user1, but got %v", foundDiary)
	}

	notFoundDiary := diaries.FindByUsername("user3")
	if notFoundDiary != nil {
		t.Errorf("Expected not to find user3, but got %v", notFoundDiary)
	}
}




func TestDiariesDelete(t *testing.T) {
	diaries := &Diaries{}
	diary1 := &Diary{username: "user1", password: "pass1"}
	diary2 := &Diary{username: "user2", password: "pass2"}

	diaries.Add(diary1)
	diaries.Add(diary2)

	diaries.Delete("user1")

	if len(diaries.diaries) != 1 {
		t.Errorf("Expected 1 diary after deletion, got %d", len(diaries.diaries))
	}

	if diaries.FindByUsername("user1") != nil {
		t.Errorf("Expected user1 to be deleted, but it was found")
	}

	if diaries.FindByUsername("user2") == nil {
		t.Errorf("Expected user2 to still exist, but it was not found")
	}

	diaries.Delete("nonexistent_user")
	if len(diaries.diaries) != 1 {
		t.Errorf("Expected 1 diary after deleting nonexistent user, got %d", len(diaries.diaries))
	}
}

