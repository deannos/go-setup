package entity

import (
	"time"
)

// User represents the core domain model
// It is framework-agnostic and contains only business logic
type User struct {
    ID        int64
    Name      string
    Email     string
    CreatedAt time.Time
    UpdatedAt time.Time
}

// CreateUserRequest is the request payload for creating a user
type CreateUserRequest struct {
    Name  string `json:"name" binding:"required,min=2,max=100"`
    Email string `json:"email" binding:"required,email"`
}

// UpdateUserRequest is the request payload for updating a user
type UpdateUserRequest struct {
    Name  string `json:"name" binding:"required,min=2,max=100"`
    Email string `json:"email" binding:"required,email"`
}

// Validate validates the user entity
// func (u *User) Validate() error {
//     if u.Name == "" {
//         return ErrUserNameRequired
//     }
//     if u.Email == "" {
//         return ErrUserEmailRequired
//     }
//     return nil
// }