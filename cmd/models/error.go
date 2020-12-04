package models

// HTTPError - custom http error type
type HTTPError struct {
	Code    int    `json:"code"`
	Key     string `json:"error"`
	Message string `json:"message"`
}

// NewHTTPError - error message format func
func NewHTTPError(code int, key string, msg string) *HTTPError {
	return &HTTPError{
		Code:    code,
		Key:     key,
		Message: msg,
	}
}

func (e *HTTPError) Error() string {
	return e.Key + ": " + e.Message
}
