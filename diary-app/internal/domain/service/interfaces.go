// internal/service/interfaces.go
package service

import (
	"diary-app/internal/api/dto"
	"diary-app/internal/domain/models"
)

type DiaryService interface {
	CreateDiary(request dto.CreateDiaryRequest) (*models.Diary, error)
	UnlockDiary(request dto.UnlockDiaryRequest) (*models.Diary, error)
	LockDiary(username string) (*models.Diary, error)
	CreateEntry(request dto.CreateEntryRequest) (*models.Diary, error)
	DeleteEntry(username string, entryID int) (*models.Diary, error)
	FindEntryByID(username string, entryID int) (*models.Entry, error)
	UpdateEntry(request dto.UpdateEntryRequest) (*models.Diary, error)
}