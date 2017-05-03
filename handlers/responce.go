package handlers

type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Data interface{}
}

func GetErrorMessage(code int, message string) ErrorResponse {
	err := &Error{Code: code, Message: message}
	return ErrorResponse{Error: *err}
}

func GetSuccessMessage(data interface{}) (SuccessResponse) {
	return SuccessResponse{}
}