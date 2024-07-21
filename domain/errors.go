package domain

import "errors"

var (
	ErrNotFoundFile     = errors.New("file not found")
	ErrNotFoundFileName = errors.New("file name not found")
	ErrBadParamFile     = errors.New("given file is not valid")

	// not yet map
	ErrFileExceedsLimit = errors.New("file size exceeds the allowed limit")
	ErrFileNotSupported = errors.New("file type is not allowed")
	ErrExistingFile     = errors.New("file already exist")
	ErrOpenAIService  = errors.New("openAI service error")

	ErrUploadBook  = errors.New("failed to upload book")
	ErrSummaryBook = errors.New("failed to summary book")

	ErrInternalServerError = errors.New("internal Server Error")
)
