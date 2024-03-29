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
	// Update the session in our DB
	// row := ss.DB.QueryRow(`
	// 	UPDATE sessions
	// 	SET token_hash = $2
	// 	WHERE user_id = $1
	// 	RETURNING id;`, session.UserID, session.TokenHash)
	// err = row.Scan(&session.ID)
	// if err == sql.ErrNoRows {
	// 	// Store the session in our DB
	// 	row = ss.DB.QueryRow(`
	// 		INSERT INTO sessions (user_id, token_hash)
	// 		VALUES ($1, $2)
	// 		RETURNING id;`, session.UserID, session.TokenHash)
	// 	err = row.Scan(&session.ID)
	// }
	// Update or store (on-conflict) the session object
	row := ss.DB.QueryRow(`
		INSERT INTO sessions (user_id, token_hash)
		VALUES ($1, $2) ON CONFLICT (user_id) DO
		UPDATE
		SET token_hash = $2
		RETURNING id;`, session.UserID, session.TokenHash)
	err = row.Scan(&session.ID)

	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// Implement SessionService.User
	// Hash the session token
	tokenHash := ss.hash(token)

	// 2. Query for the session w/ that hash
	// var user User
	// row := ss.DB.QueryRow(`
	// SELECT user_id
	// FROM sessions
	// where TOKEN_HASH = $1;`, tokenHash)
	// err := row.Scan(&user.ID)
	// if err != nil {
	// 	return nil, fmt.Errorf("user: %w", err)
	// }
	// // 3. Using the UserID from session, query for that user
	// row = ss.DB.QueryRow(`
	// SELECT email, password_hash
	// FROM users WHERE id = $1;`, user.ID)
	// err = row.Scan(&user.Email, &user.PasswordHash)
	// if err != nil {
	// 	return nil, fmt.Errorf("user: %w", err)
	// }

	// Slimmed down version 2 & 3 consolidated with join...
	var user User
	row := ss.DB.QueryRow(`
	SELECT user_id,
	  users.email,
	  users.password_hash
	FROM sessions
	  JOIN users ON users.id = sessions.user_id
	where sessions.token_hash = $1;`, tokenHash)
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	// 4. return the user
	return &user, nil

	// return nil, nil
}

func (ss *SessionService) Delete(token string) error {
	tokenHash := ss.hash(token)
	_, err := ss.DB.Exec(`
		DELETE FROM sessions
		WHERE token_hash = $1;`, tokenHash)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	// base64 encode the data to string
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}
