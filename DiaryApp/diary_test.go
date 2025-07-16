package main

import (
	"testing"
)

func TestNewDiary(t *testing.T) {
	username := "testuser"
	password := "testpass"
	diary := NewDiary(username, password)

	if diary.username != username {
		t.Errorf("Expected username %s, got %s", username, diary.username)
	}
	if diary.password != password {
		t.Errorf("Expected password %s, got %s", password, diary.password)
	}
	if diary.isLocked != true {
		t.Errorf("Expected isLocked to be true, got %t", diary.isLocked)
	}
	if diary.entries == nil {
		t.Errorf("Expected entries to be initialized, got nil")
	}
}




func TestDiary_UnlockDiary(t *testing.T) {
	diary := NewDiary("testuser", "testpass")
	diary.UnlockDiary("testpass")

	if diary.isLocked != false {
		t.Errorf("Expected isLocked to be false after unlocking, got %t", diary.isLocked)
	}

	diary.UnlockDiary("wrongpass")
	if diary.isLocked != false {
		t.Errorf("Expected isLocked to remain false after wrong password, got %t", diary.isLocked)
	}
}




func TestDiary_LockDiary(t *testing.T) {
	diary := NewDiary("testuser", "testpass")
	diary.UnlockDiary("testpass") // Unlock first to test locking

	diary.LockDiary()

	if diary.isLocked != true {
		t.Errorf("Expected isLocked to be true after locking, got %t", diary.isLocked)
	}
}




func TestDiaryIsLocked(t *testing.T) {
	diary := NewDiary("testuser", "testpass")

	if diary.IsLocked() != true {
		t.Errorf("Expected IsLocked to be true initially, got %t", diary.IsLocked())
	}

	diary.UnlockDiary("testpass")
	if diary.IsLocked() != false {
		t.Errorf("Expected IsLocked to be false after unlocking, got %t", diary.IsLocked())
	}

	diary.LockDiary()
	if diary.IsLocked() != true {
		t.Errorf("Expected IsLocked to be true after locking, got %t", diary.IsLocked())
	}
}




func TestDiaryCreateEntry(t *testing.T) {
	diary := NewDiary("testuser", "testpass")
	diary.UnlockDiary("testpass")

	diary.CreateEntry("Title 1", "Body 1")

	if len(diary.entries) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(diary.entries))
	}

	if diary.entries[0].Title != "Title 1" || diary.entries[0].Body != "Body 1" {
		t.Errorf("Expected entry with Title 1 and Body 1, got %s, %s", diary.entries[0].Title, diary.entries[0].Body)
	}

	diary.LockDiary()
	diary.CreateEntry("Title 2", "Body 2")

	if len(diary.entries) != 1 {
		t.Errorf("Expected 1 entry after trying to create while locked, got %d", len(diary.entries))
	}
}




func TestDiaryDeleteEntry(t *testing.T) {
	diary := NewDiary("testuser", "testpass")
	diary.UnlockDiary("testpass")

	diary.CreateEntry("Title 1", "Body 1")
	diary.CreateEntry("Title 2", "Body 2")

	diary.DeleteEntry(1)

	if len(diary.entries) != 1 {
		t.Errorf("Expected 1 entry after deletion, got %d", len(diary.entries))
	}

	if diary.entries[0].ID != 2 {
		t.Errorf("Expected remaining entry ID to be 2, got %d", diary.entries[0].ID)
	}

	diary.DeleteEntry(99) 
	if len(diary.entries) != 1 {
		t.Errorf("Expected 1 entry after deleting non-existent entry, got %d", len(diary.entries))
	}

	diary.LockDiary()
	diary.DeleteEntry(2)

	if len(diary.entries) != 1 {
		t.Errorf("Expected 1 entry after trying to delete while locked, got %d", len(diary.entries))
	}
}




func TestDiaryFindEntryByID(t *testing.T) {
	diary := NewDiary("testuser", "testpass")
	diary.UnlockDiary("testpass")

	diary.CreateEntry("Title 1", "Body 1")
	diary.CreateEntry("Title 2", "Body 2")

	foundEntry := diary.FindEntryByID(1)
	if foundEntry == nil || foundEntry.ID != 1 {
		t.Errorf("Expected to find entry with ID 1, but got %v", foundEntry)
	}

	notFoundEntry := diary.FindEntryByID(99)
	if notFoundEntry != nil {
		t.Errorf("Expected not to find entry with ID 99, but got %v", notFoundEntry)
	}

	diary.LockDiary()
	lockedFoundEntry := diary.FindEntryByID(1)
	if lockedFoundEntry != nil {
		t.Errorf("Expected not to find entry when locked, but got %v", lockedFoundEntry)
	}
}




func TestDiary_UpdateEntry(t *testing.T) {
	diary := NewDiary("testuser", "testpass")
	diary.UnlockDiary("testpass")

	diary.CreateEntry("Title 1", "Body 1")

	diary.UpdateEntry(1, "Updated Title", "Updated Body")

	updatedEntry := diary.FindEntryByID(1)
	if updatedEntry == nil || updatedEntry.Title != "Updated Title" || updatedEntry.Body != "Updated Body" {
		t.Errorf("Expected entry to be updated, got %v", updatedEntry)
	}

	diary.UpdateEntry(99, "NonExistent Title", "NonExistent Body") // Update non-existent entry
	// No change expected, so FindEntryByID(99) should still return nil
	if diary.FindEntryByID(99) != nil {
		t.Errorf("Expected no entry with ID 99, but found one after attempted update")
	}

	diary.LockDiary()
	diary.UpdateEntry(1, "Locked Title", "Locked Body")

	lockedEntry := diary.FindEntryByID(1)
	if lockedEntry != nil && lockedEntry.Title == "Locked Title" {
		t.Errorf("Expected entry not to be updated when locked, but it was")
	}
}
