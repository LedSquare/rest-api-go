package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOk    = "Success"
	StatusError = "Error"
)

func Success() Response {
	return Response{
		Status: StatusOk,
	}
}

func Error(message string) Response {
	return Response{
		Status: StatusError,
		Error:  message,
	}
}

func ValidaationErrors(errors validator.ValidationErrors) Response {
	var errorMessages []string

	for _, err := range errors {
		switch err.ActualTag() {
		case "required":
			errorMessages = append(errorMessages, fmt.Sprintf("field %s is a required field", err.Field()))
		case "url":
			errorMessages = append(errorMessages, fmt.Sprintf("field %s is a required field", err.Field()))
		default:
			errorMessages = append(errorMessages, fmt.Sprintf("field %s is a required field", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errorMessages, ", "),
	}
}
