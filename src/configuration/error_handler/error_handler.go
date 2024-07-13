package errorHandler

import "net/http"

type ErrorHandler struct {
	Message string `json:"message"`
	Err string `json:"error"`
	Code int `json:"code"`
	Causes []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

func (e *ErrorHandler) Error() string {	
	return e.Message
}

func NewErrorHandler(message string, err string, code int, causes []Causes) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err: err,
		Code: code,
		Causes: causes,
	}
}

func NewBadRequestError(message string) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err: "Bad Request",
		Code: http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err: "Bad Request",
		Code: http.StatusBadRequest,
		Causes: causes,
	}
}

func NewUnauthorizedError(message string) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err: "Unauthorized",
		Code: http.StatusUnauthorized,
	}
}

func NewNotFoundError(message string) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err: "Not Found",
		Code: http.StatusNotFound,
	}
}	

func NewInternalServerError(message string) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err: "Internal Server Error",
		Code: http.StatusInternalServerError,
	}
}

func NewForbiddenError(message string) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err: "Forbidden",
		Code: http.StatusForbidden,
	}
}
