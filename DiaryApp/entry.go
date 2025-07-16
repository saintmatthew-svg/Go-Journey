package main

import (
	"time"
)

type Entry struct {
	ID          int
	Title       string
	Body        string
	DateCreated time.Time
}

func NewEntry(id int, title, body string) *Entry {
	return &Entry{
		ID:          id,
		Title:       title,
		Body:        body,
		DateCreated: time.Now(),
	}
}
