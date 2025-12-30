package errors

import "fmt"

type AppError struct {
    Message string
    Code    int
}

func (e *AppError) Error() string {
    return e.Message
}

var (
    ErrNotFound       = &AppError{"Resource not found", 404}
    ErrValidation     = &AppError{"Validation failed", 400}
    ErrUnauthorized   = &AppError{"Unauthorized", 401}
    ErrForbidden      = &AppError{"Forbidden", 403}
    ErrConflict       = &AppError{"Conflict", 409}
    ErrInternalServer = &AppError{"Internal server error", 500}
    
    ErrUserNameRequired  = &AppError{"User name is required", 400}
    ErrUserEmailRequired = &AppError{"User email is required", 400}
)

func NewError(message string, code int) error {
    return &AppError{message, code}
}

func NewDatabaseError(err error) error {
    return &AppError{
        Message: fmt.Sprintf("Database error: %v", err),
        Code:    500,
    }
}

func ToHTTP(err error) (int, string) {
    if appErr, ok := err.(*AppError); ok {
        return appErr.Code, appErr.Message
    }
    return 500, "Internal server error"
}