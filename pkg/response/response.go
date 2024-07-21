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

	domain.ErrFileExceedsLimit: "File size exceeds the allowed limit",
	domain.ErrFileNotSupported: "File type is not allowed",
	domain.ErrExistingFile:     "File already exist",
	domain.ErrOpenAIService:  "OpenAI service error",

	domain.ErrUploadBook:       "Failed to upload book",
	domain.ErrSummaryBook:      "Failed to summary book",

	domain.ErrInternalServerError: "Internal Server Error",
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
	case domain.ErrFileExceedsLimit:
		return http.StatusBadRequest
	case domain.ErrFileNotSupported:
		return http.StatusBadRequest
	case domain.ErrExistingFile:
		return http.StatusConflict
	case domain.ErrOpenAIService:
		return http.StatusInternalServerError
	case domain.ErrUploadBook:
		return http.StatusInternalServerError
	case domain.ErrSummaryBook:
		return http.StatusInternalServerError
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func GetError(err error) ResponseError {
	return ResponseError{Error: err, Message: getErrorMessage(err)}
}
