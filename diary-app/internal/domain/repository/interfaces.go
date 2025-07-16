// internal/repository/interfaces.go
package repository

import "diary-app/internal/domain/models"

type DiaryRepository interface {
    Add(username, password string) (*models.Diary, error)
    FindByUsername(username string) (*models.Diary, error)
    Delete(username, password string) error
    GetAll() *models.Diaries
    UpdatePassword(username, oldPassword, newPassword string) error
}