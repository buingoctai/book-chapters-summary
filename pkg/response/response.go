package response

import (
	"github.com/buingoctai/book-chapters-summary/domain"
	"net/http"
)

type ResponseError struct {
	Error   error
	Message string
}

var ErrorMessageMap = map[error]string{
	domain.ErrNotFoundFile:     "File not found",
	domain.ErrNotFoundFileName: "File name not found",
	domain.ErrBadParamFile:     "Given Param is not valid",
	domain.ErrUploadBook:       "Failed to upload book",
	domain.ErrSummaryBook:      "Failed to summary book",

	domain.ErrInternalServerError: "Internal Server Error",
	domain.ErrNotFound:            "Not Found",
	domain.ErrConflict:            "Conflict",
}

func getErrorMessage(err error) string {
	if msg, ok := ErrorMessageMap[err]; ok {
		return msg
	}
	return "Internal Server Error"
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrNotFoundFile:
		return http.StatusBadRequest
	case domain.ErrNotFoundFileName:
		return http.StatusBadRequest
	case domain.ErrBadParamFile:
		return http.StatusBadRequest
	case domain.ErrUploadBook:
		return http.StatusInternalServerError
	case domain.ErrSummaryBook:
		return http.StatusInternalServerError
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func GetError(err error) ResponseError {
	return ResponseError{Error: err, Message: getErrorMessage(err)}
}
