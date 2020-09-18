package middleware

import "errors"

// Throw401 is a shortcut to throw HTTP:500 Errors
func Throw401(Message string) *HTTPError {
	return &HTTPError{
		Err:     errors.New("service access error"),
		Message: Message,
		Status:  401,
	}
}

// Throw404 is a shortcut to throw HTTP:404 Errors
func Throw404(Message string) *HTTPError {
	return &HTTPError{
		Err:     errors.New("database error"),
		Message: Message,
		Status:  404,
	}
}

// Throw500 is a shortcut to throw HTTP:500 Errors
func Throw500(Message string) *HTTPError {
	return &HTTPError{
		Err:     errors.New("general server error"),
		Message: Message,
		Status:  500,
	}
}
