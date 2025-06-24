package lib

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Error struct {
	// false => Error originating from binance's side
	// true => Originating from the Binance-Go library
	IsLocalError bool

	StatusCode int

	Code int

	Message string
}

// Implement the `Error` method to satisfy the `error` interface
func (e *Error) Error() string {
	str := "[BINANCE ERROR] StatusCode " + strconv.Itoa(e.StatusCode)
	if e.IsLocalError {
		str = "[LIB ERROR]"
	}
	return fmt.Sprintf("%s - Code %d: \"%s\"", str, e.Code, e.Message)
}

func newError(isLocal bool, statusCode int, code int, message string) *Error {
	err := &Error{
		IsLocalError: isLocal,
		StatusCode:   statusCode,
		Code:         code,
		Message:      message,
	}

	return err
}

////

func LocalError(code int, msg string) *Error {
	return newError(true, 0, code, msg)
}

type binanceErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//

// Processes an erroneous 4XX HTTP Response
// Returns the library Error type
// In the case of an error parsing the error body, it returns a secondaly unmarshall error
func BinanceError(statusCode int, body []byte) (BinanceError *Error, UnmarshallError error) {
	var errResponse binanceErrorResponse

	unmarshallErr := json.Unmarshal(body, &errResponse)
	if unmarshallErr != nil {
		return nil, unmarshallErr
	}

	err := newError(false, statusCode, errResponse.Code, errResponse.Msg)

	return err, nil
}
