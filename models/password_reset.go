package models

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	// DefaultResetDuration is the default time that
	// a PasswordReset is valid for.
	DefaultResetDuration = 1 * time.Hour
)

type PasswordReset struct {
	ID     int
	UserID int
	// Token is only set when a PasswordReset is created
	Token     string
	TokenHash string
	ExpiresAt time.Time
}

type PasswordResetService struct {
	DB *sql.DB
	// Size of tokens. If not set, MinBytesPerToken used
	BytesPerToken int
	// Duration is amont of time pw reset is valid for
	Duration time.Duration
}

func (service *PasswordResetService) Create(email string) (*PasswordReset, error) {
	return nil, fmt.Errorf("TODO: Implement PasswordResetService.Create")
}

func (service *PasswordResetService) Consume(token string) (*User, error) {
	return nil, fmt.Errorf("TODO: Implement PasswordResetService.Consume")
}
