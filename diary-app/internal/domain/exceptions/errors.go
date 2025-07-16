// internal/domain/exceptions/errors.go
package exceptions

import "errors"

var (
	ErrDiaryAlreadyExists = errors.New("diary with this username already exists")
	ErrDiaryNotFound      = errors.New("diary not found")
	ErrInvalidPassword    = errors.New("invalid userame or password")
	ErrDiaryLocked        = errors.New("diary is locked, login to get access")
	ErrEntryNotFound      = errors.New("entry not found")
)