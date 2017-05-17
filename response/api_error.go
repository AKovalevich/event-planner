package response

import (
	"github.com/AKovalevich/event-planner/app"
)

// APIError represents an error that can be sent in an error response.
type APIError struct {
	// Status represents the HTTP status code
	Status int `json:"-"`
	// ErrorCode is the code uniquely identifying an error
	ErrorCode string `json:"error_code"`
	// Message is the error message that may be displayed to end users
	Message string `json:"message"`
	// DeveloperMessage is the error message that is mainly meant for developers
	DeveloperMessage string `json:"developer_message,omitempty"`
	// Details specifies the additional error information
	Details interface{} `json:"details,omitempty"`
}

// Error returns the error message.
func (e APIError) Error() string {
	return e.Message
}

// StatusCode returns the HTTP status code.
func (e APIError) StatusCode() int {
	return e.Status
}

func (e APIError) Struct() map[string]interface{} {
	response := map[string]interface{}{}

	response["Status"] = e.Status
	response["Code"] = e.ErrorCode
	response["Message"] = e.Message

	if app.Config().Debug {
		response["DeveloperMessage"] = e.DeveloperMessage
		response["Details"] = e.Details
	}

	return response
}
