package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"github.com/buingoctai/book-chapters-summary/pkg/contexts"
)

func SplitIntoChapters(content string, separatorPattern string) []string {
    templateType := "word"
    adaptor := context.NewContentAdaptor(templateType)
    chapters := adaptor.AdaptContent(content)

    return chapters
}

func GetLocalPath(fileName string) string {
	path := os.Getenv("LOCAL_FILE_PATH")
	if path == "" {
		path = "./"
	}

	return path + fileName
}


func GetTempsFolderPath() string {
	return os.Getenv("TEMP_FILE_PATH")
}

func CalculateChecksum(fileBytes []byte) string {
	hash := sha256.Sum256(fileBytes)
	return hex.EncodeToString(hash[:])
}

func IsFileExistsInFolder(folderPath string, checksum string) bool {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return false
	}
	for _, file := range files {
		if file.Name() == checksum {
			return true
		}
	}
	return false
}