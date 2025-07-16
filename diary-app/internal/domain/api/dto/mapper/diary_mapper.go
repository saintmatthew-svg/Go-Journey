// internal/api/mapper/diary_mapper.go
package mapper

import (
	"diary-app/internal/api/dto"
	"diary-app/internal/domain/models"
)

func ToDiaryResponse(diary *models.Diary, includeEntries bool) dto.DiaryResponse {
	response := dto.DiaryResponse{
		Username: diary.Username,
		IsLocked: diary.IsLocked,
	}
	
	if includeEntries {
		response.Entries = make([]dto.EntryResponse, len(diary.Entries))
		for i, entry := range diary.Entries {
			response.Entries[i] = ToEntryResponse(entry)
		}
	}
	
	return response
}

func ToEntryResponse(entry models.Entry) dto.EntryResponse {
	return dto.EntryResponse{
		ID:          entry.ID,
		Title:       entry.Title,
		Body:        entry.Body,
		DateCreated: entry.DateCreated,
	}
}