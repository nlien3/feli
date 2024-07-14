package services

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"regexp"

	"backend/clients"
)

// Username validation
func validateUsername(username string) error {
	if len(username) < 3 || len(username) > 20 {
		return errors.New("username must be between 3 and 20 characters")
	}
	// Ensure username only contains valid characters
	validUsername := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !validUsername.MatchString(username) {
		return errors.New("username can only contain letters, numbers, and underscores")
	}
	return nil
}

// Password validation
func validatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	return nil
}

func hashPasswordMD5(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func RegisterS(username, password, tipo string) error {
	// Validate username
	if err := validateUsername(username); err != nil {
		return err
	}

	// Validate password
	if err := validatePassword(password); err != nil {
		return err
	}

	// Check if the user already exists
	err := clients.SearchUser(username)
	if err == nil {
		return errors.New("username already taken")
	}

	// Hash the password using MD5
	hashedPassword := hashPasswordMD5(password)

	// Create the new user
	err = clients.CreateUser(username, hashedPassword, tipo)
	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}
