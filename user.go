// Package models contains weaver-ci model structs
package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// User model
type User struct {
	UserID       uuid.UUID
	EmailAddress string
	LastLogin    *time.Time
	ModifiedOn   *time.Time
	CreatedOn    *time.Time
}

// GithubLogin model
type GithubLogin struct {
	UserID     uuid.UUID
	ResetToken string
}

// AssemblaLogin model
type AssemblaLogin struct {
	UserID     uuid.UUID
	ResetToken string
}
