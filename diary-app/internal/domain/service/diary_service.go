// internal/service/diary_service.go
package service

import (
	"diary-app/internal/api/dto"
	"diary-app/internal/domain/exceptions"
	"diary-app/internal/domain/models"
	"diary-app/internal/repository"
)

type diaryService struct {
	diaryRepo repository.DiaryRepository
}

func NewDiaryService(diaryRepo repository.DiaryRepository) DiaryService {
	return &diaryService{diaryRepo: diaryRepo}
}

func (s *diaryService) CreateDiary(request dto.CreateDiaryRequest) (*models.Diary, error) {
	return s.diaryRepo.Add(request.Username, request.Password)
}

func (s *diaryService) UnlockDiary(request dto.UnlockDiaryRequest) (*models.Diary, error) {
	diary, err := s.diaryRepo.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}
	
	if diary.Password != request.Password {
		return nil, exceptions.ErrInvalidPassword
	}
	
	diary.IsLocked = false
	if err := s.diaryRepo.Update(diary); err != nil {
		return nil, err
	}
	
	return diary, nil
}

func (s *diaryService) LockDiary(username string) (*models.Diary, error) {
	diary, err := s.diaryRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	
	diary.IsLocked = true
	if err := s.diaryRepo.Update(diary); err != nil {
		return nil, err
	}
	
	return diary, nil
}

func (s *diaryService) CreateEntry(request dto.CreateEntryRequest) (*models.Diary, error) {
	diary, err := s.diaryRepo.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}
	
	if diary.IsLocked {
		return nil, exceptions.ErrDiaryLocked
	}
	
	diary.CreateEntry(request.Title, request.Body)
	if err := s.diaryRepo.Update(diary); err != nil {
		return nil, err
	}
	
	return diary, nil
}

// Implement other service methods following the same pattern...