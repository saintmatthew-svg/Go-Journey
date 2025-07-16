// internal/repository/diary_repo.go
package repository

import (
    "diary-app/internal/domain/exceptions"
    "diary-app/internal/domain/models"
    "sync"
)

type diaryRepository struct {
    diaries *models.Diaries
    mu      sync.Mutex
}

func NewDiaryRepository() DiaryRepository {
    return &diaryRepository{
        diaries: models.NewDiaries(),
    }
}

func (r *diaryRepository) Add(username, password string) (*models.Diary, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    if existing := r.diaries.FindByUsername(username); existing != nil {
        return nil, exceptions.ErrDiaryAlreadyExists
    }
    
    diary := models.NewDiary(username, password)
    r.diaries.Add(diary)
    return diary, nil
}

func (r *diaryRepository) FindByUsername(username string) (*models.Diary, error) {
    if diary := r.diaries.FindByUsername(username); diary != nil {
        return diary, nil
    }
    return nil, exceptions.ErrDiaryNotFound
}

func (r *diaryRepository) Delete(username, password string) error {
    if !r.diaries.Delete(username, password) {
        return exceptions.ErrDiaryNotFound
    }
    return nil
}

func (r *diaryRepository) GetAll() *models.Diaries {
    return r.diaries
}