package berror

import (
	"errors"
	"fmt"
)

const (
	networkErrorCode      = 1
	responseReadErrorCode = 2
	parsingErrorCode      = 3
	validationErrorCode   = 4
	signatureErrorCode    = 5
	invalidValueErrorCode = 6
	notFoundErrorCode     = 7
	timeoutErrorCode      = 8
)

type bError struct {
	islocal bool
	code    int
	message string
	error   error
}

func (e bError) Error() string {
	if e.islocal {
		return fmt.Sprintf("Local error %d \"%s\"", e.code, e.message)
	}
	return fmt.Sprintf("Binance API error %d \"%s\"", e.code, e.message)
}

func (e bError) IsLocal() bool {
	return e.islocal
}

func (e bError) Code() int {
	return e.code
}

func (e bError) Message() string {
	return e.message
}

func (e bError) Matches(errCode ErrorCode) bool {
	return e.code == errCode.Code
}

func (e bError) Unwrap() error {
	return e.error
}

// //
// apiError represents an error returned by the Binance API.
// //
func NewAPIError(code int, message string) Error {
	return bError{
		islocal: false,
		code:    code,
		message: message,
		error:   nil,
	}
}

// //
// NewResponseReadError represents an error that occurred while reading the response from the Binance API.
// //
func NewResponseReadError(err error) Error {
	return bError{
		islocal: true,
		code:    responseReadErrorCode,
		message: err.Error(),
		error:   err,
	}
}

// //
// networkError represents an error that occurred due to network issues or other local problems.
// //
func NewNetworkError(err error) Error {
	return bError{
		islocal: true,
		code:    networkErrorCode,
		message: err.Error(),
		error:   err,
	}
}

// //
// parsingError represents an error that occurred due to parsing issues.
// //
func NewParseError(err error) Error {
	return bError{
		islocal: true,
		code:    parsingErrorCode,
		message: err.Error(),
		error:   err,
	}
}

// //
// validationError represents an error that occurred due to validation issues.
// //
func NewValidationError(err error) Error {
	return bError{
		islocal: true,
		code:    validationErrorCode,
		message: err.Error(),
		error:   err,
	}
}

// //
// signatureError represents an error that occurred due to signature issues.
// //
func NewSignatureError(err error) Error {
	return bError{
		islocal: true,
		code:    signatureErrorCode,
		message: err.Error(),
		error:   err,
	}
}

// //
//
// //
func NewInvalidValueError(message string) Error {
	return bError{
		islocal: true,
		code:    invalidValueErrorCode,
		message: message,
		error:   errors.New(message),
	}
}

// //
//
// //
func NewNotFoundError(message string) Error {
	return bError{
		islocal: true,
		code:    notFoundErrorCode,
		message: message,
		error:   errors.New(message),
	}
}

// //
//
// //
func NewTimeoutError(message string) Error {
	return bError{
		islocal: true,
		code:    timeoutErrorCode,
		message: message,
		error:   errors.New(message),
	}
}
