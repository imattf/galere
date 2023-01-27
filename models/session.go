package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/imattf/galere/rand"
)

const (
	// The minimum number of bytes to be used for each session
	MinBytesPerToken = 32
)

type Session struct {
	ID        int
	UserID    int
	Token     string //Only set on new session
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
	// BytesPerToken is used to determine how many bytes to use when generating
	// each session token. If this value is not set or is less than the
	// MinBytesPerToken const it will be ignored and MinBytesPerToken will be
	// used.
	BytesPerToken int
}

// Create new session for the user provided.
func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	// Create the session token
	// token, err := rand.SessionToken()
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	// Implement SessionService.Create
	session := Session{
		UserID: userID,
		Token:  token,
		// Set the token hash
		TokenHash: ss.hash(token),
	}
	// TODO: Store the session in our DB
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	return nil, nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	// base64 encode the data to string
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}
