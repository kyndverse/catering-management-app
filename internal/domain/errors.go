package domain

type ErrorType string

const (
	ErrorNotFound        ErrorType = "NOT_FOUND"
	ErrorConflict        ErrorType = "CONFLICT"
	ErrorUnauthorized    ErrorType = "UNAUTHORIZED"
	ErrorForbidden       ErrorType = "FORBIDDEN"
	ErrorBadRequest      ErrorType = "BAD_REQUEST"
	ErrInvalidCredential ErrorType = "UNAUTHORIZED"
)

type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewError(errorType ErrorType, message string) *AppError {
	return &AppError{
		Type:    errorType,
		Message: message,
	}
}

func WrapError(errorType ErrorType, message string, err error) *AppError {
	return &AppError{
		Type:    errorType,
		Message: message,
		Err:     err,
	}
}
