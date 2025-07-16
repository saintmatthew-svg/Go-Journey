// internal/api/dto/request.go
package dto

type CreateDiaryRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UnlockDiaryRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateEntryRequest struct {
	Username string `json:"username" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
}

type UpdateEntryRequest struct {
	Username string `json:"username" validate:"required"`
	EntryID  int    `json:"entry_id" validate:"required"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}