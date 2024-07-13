package validator

import (
	"mime/multipart"
	"path/filepath"
	"strings"
	"github.com/buingoctai/book-chapters-summary/domain"
)

const maxSize = 10485760 // 10MB
var allowedTypes = map[string]bool{
	"jpg":  true,
	"jpeg": true,
	"png":  true,
}

func IsBookValid(formFile *multipart.FileHeader) (bool, error) {
	//validate file name
	if formFile.Filename == "" {
		return false, domain.ErrNotFoundFileName
	}
	// Validate file size
	if formFile.Size > maxSize {
		return false, domain.ErrFileExceedsLimit
	}

	// Validate file type
	ext := strings.TrimPrefix(filepath.Ext(formFile.Filename), ".")
	if !allowedTypes[strings.ToLower(ext)] {
		return false, domain.ErrFileNotSupported
	}

	return true, nil
}