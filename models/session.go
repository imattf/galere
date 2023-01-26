package models

import (
	"database/sql"
	"fmt"

	"github.com/imattf/galere/rand"
)

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
	// Create the session token
	token, err := rand.SessionToken()
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	// TODO: Hash the session token
	// Implement SessionService.Create
	session := Session{
		UserID: userID,
		Token:  token,
		// TODO: Set the token hash
	}
	// TODO: Store the session in our DB
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	return nil, nil
}
