package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/imattf/galere/rand"
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
	// Verify we have a valid email address for a user and get that user's ID
	email = strings.ToLower(email)
	var userID int
	row := service.DB.QueryRow(`
		SELECT id FROM users WHERE email = $1;`, email)
	err := row.Scan(&userID)
	if err != nil {
		// TODO: Consider returning a specific error when the user does not exist
		return nil, fmt.Errorf("create: find email: %w", err)
	}

	// fmt.Println("Create: email: ", email) //DEBUG

	// Build the PasswordReset
	bytesPerToken := service.BytesPerToken
	if bytesPerToken == 0 {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: build pwr: %w", err)
	}

	// fmt.Println("Create: token: ", token) //DEBUG

	duration := service.Duration
	if duration == 0 {
		duration = DefaultResetDuration
	}

	// Construct password reset object
	pwReset := PasswordReset{
		UserID:    userID,
		Token:     token,
		TokenHash: service.hash(token),
		ExpiresAt: time.Now().Add(duration),
	}

	// Insert the PasswordReset into the DB
	row = service.DB.QueryRow(`
	INSERT INTO password_resets (user_id, token_hash, expires_at)
	VALUES ($1, $2, $3) ON CONFLICT (user_id) DO
	UPDATE
	SET token_hash = $2, expires_at = $3
	RETURNING id;`, pwReset.UserID, pwReset.TokenHash, pwReset.ExpiresAt)
	err = row.Scan(&pwReset.ID)
	if err != nil {
		return nil, fmt.Errorf("create: insert:pw-reset %w", err)
	}
	// fmt.Println("Storing: ", pwReset) //DEBUG
	return &pwReset, nil
}

func (service *PasswordResetService) Consume(token string) (*User, error) {
	// make sure provided token is valid & not expired (stored password reset)
	// make sure we have user info
	// (use join to find user & password)
	// delete the now used token (password reset)

	// match provided token against what user responed to
	tokenHash := service.hash(token)

	// fmt.Println("Token     in Consume: ", token)     //DEBUG
	// fmt.Println("TokenHash in Consume: ", tokenHash) //DEBUG
	var user User
	var pwReset PasswordReset
	row := service.DB.QueryRow(`
	  SELECT password_resets.id,
	    password_resets.expires_at,
		users.id,
		users.email,
		users.password_hash
	  FROM password_resets
	    JOIN users ON users.id = password_resets.user_id
	  WHERE password_resets.token_hash = $1;`, tokenHash)
	err := row.Scan(
		&pwReset.ID, &pwReset.ExpiresAt,
		&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		fmt.Println("tokenHash = ", tokenHash)
		return nil, fmt.Errorf("consume: find pw-reset by tokenhash: %w", err)
	}

	// Validate pwReset has not expired
	if time.Now().After(pwReset.ExpiresAt) {
		return nil, fmt.Errorf("token expired: %v", token)
	}

	// Delete the used password reset token
	err = service.delete(pwReset.ID)
	if err != nil {
		return nil, fmt.Errorf("consume delete pw-reset: %w", err)
	}
	return &user, nil
}

func (service *PasswordResetService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}

func (service *PasswordResetService) delete(id int) error {
	_, err := service.DB.Exec(`
	  DELETE FROM password_resets
	  WHERE id = $1;`, id)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}
