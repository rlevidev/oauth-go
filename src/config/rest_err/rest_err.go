package resterr

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Status  int      `json:"status"`
	Err     string   `json:"error"`
	Causes  []Causes `json:"causes,omitempty"`
}

func (r *RestErr) Error() string {
	return r.Message
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewRestErr(message string, status int, err string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Err:     err,
		Causes: causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Err:     "bad_request",
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Err:     "bad_request",
		Causes:  causes,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Err:     "not_found",
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusUnauthorized,
		Err:     "unauthorized",
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusForbidden,
		Err:     "forbidden",
	}
}
