package models

import "database/sql"

type Session struct {
	ID        int
	UserID    int
	Token     string //Only set on new session
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

// Create new session for the user provided.
func (ss *SessionService) Create(userID int) (*Session, error) {
	// TODO: Create the session token
	// TODO: Implement SessionService.Create
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	return nil, nil
}
