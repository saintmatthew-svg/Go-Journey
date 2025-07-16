// internal/api/dto/response.go
package dto

import "time"

type DiaryResponse struct {
	Username string          `json:"username"`
	IsLocked bool            `json:"is_locked"`
	Entries  []EntryResponse `json:"entries,omitempty"`
}

type EntryResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	DateCreated time.Time `json:"date_created"`
}

type OperationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}