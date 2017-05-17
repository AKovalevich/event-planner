package response

import (
	"net/http"
)

type validationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// InternalServerError creates a new API error representing an internal server error (HTTP 500)
func InternalServerError(err string, details string) *APIError {
	return NewAPIError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", Params{"error": err}, details)
}

// NotFound creates a new API error representing a resource-not-found error (HTTP 404)
func NotFound(resource string) *APIError {
	return NewAPIError(http.StatusNotFound, "NOT_FOUND", Params{"resource": resource}, "")
}

// Unauthorized creates a new API error representing an authentication failure (HTTP 401)
func Unauthorized(err string) *APIError {
	return NewAPIError(http.StatusUnauthorized, "UNAUTHORIZED", Params{"error": err}, "")
}

// InvalidData converts a data validation error into an API error (HTTP 400)
//func InvalidData(errs []interface{}) *APIError {
//	result := []validationError{}
//	fields := []string{}
//	for err := range errs {
//		fields = append(fields, err.Error())
//	}
//	sort.Strings(fields)
//	for _, field := range fields {
//		err := errs[field]
//		result = append(result, validationError{
//			Field: field,
//			Error: err.Error(),
//		})
//	}
//
//	err := NewAPIError(http.StatusBadRequest, "INVALID_DATA", nil)
//	err.Details = result
//
//	return err
//}

// InvalidData converts a data validation error into an API error (HTTP 400)
func InvalidData(error string) *APIError {
	err := NewAPIError(http.StatusBadRequest, "INVALID_DATA", nil, "")
	err.Details = error

	return err
}
