// internal/domain/models/diary.go
package models

import "time"

type Diary struct {
    Username string
    Password string
    IsLocked bool
    Entries  []Entry
}

func NewDiary(username, password string) *Diary {
    return &Diary{
        Username: username,
        Password: password,
        IsLocked: false,
        Entries:  make([]Entry, 0),
    }
}

func (d *Diary) Unlock(password string) bool {
    if d.Password == password {
        d.IsLocked = false
        return true
    }
    return false
}

func (d *Diary) Lock() {
    d.IsLocked = true
}

// Other Diary methods...