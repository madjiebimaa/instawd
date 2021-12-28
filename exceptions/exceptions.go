package exceptions

import "errors"

var (
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrInternalServerError   = errors.New("internal server error")
	ErrEmptyInput            = errors.New("empty input")
	ErrInvalidInput          = errors.New("invalid input")
	ErrValidationFailed      = errors.New("validation failed")
	ErrUserNameAlreadyExists = errors.New("user_name already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrUserNotFound          = errors.New("user not found")
	ErrUserNotAuthenticated  = errors.New("user not authenticated")
	ErrUserNotHavePermission = errors.New("user not have permission")
	ErrQuoteNotFound         = errors.New("quote not found")
	ErrQuoteIdAlreadyExists  = errors.New("quote id already exist")
	ErrQuoteTagNotFound      = errors.New("quote tag not found")
	ErrQuoteTagAlreadyExist  = errors.New("quote tag already exist")
	ErrAuthorNotFound        = errors.New("author not found")
)
