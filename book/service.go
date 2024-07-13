package book

import (
	"bytes"
	"github.com/buingoctai/book-chapters-summary/domain"
	"github.com/buingoctai/book-chapters-summary/pkg/third_party/openai"
	"github.com/buingoctai/book-chapters-summary/pkg/utils"
	"io"
	"mime/multipart"
	"os"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ReadUploadedFile(file multipart.File) ([]byte, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, domain.ErrBadParamFile
	}
	return content, nil
}

func (s *Service) UploadFile(fileBytes []byte, fileName string) (string, error) {
	localPath := utils.GetLocalPath(fileName)
	dst, err := os.Create(localPath)
	if err != nil {
		return localPath, domain.ErrUploadBook
	}
	defer dst.Close()

	// check duplicate file in temps folder by check same checksum
	tempsPath := utils.GetTempsFolderPath()
	checksum := utils.CalculateChecksum(fileBytes)
	if utils.IsFileExistsInFolder(tempsPath, checksum) {
		return localPath, domain.ErrExistingFile
	}

	_, err = io.Copy(dst, bytes.NewReader(fileBytes))
	if err != nil {
		return localPath, domain.ErrUploadBook
	}

	return localPath, nil
}

func (s *Service) LoadFile(fileName string) (string, error) {
	localPath := utils.GetLocalPath(fileName)
	f, err := os.Open(localPath)
	if err != nil {
		return "", domain.ErrBadParamFile
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return "", domain.ErrBadParamFile
	}
	return string(content), nil
}

func (s *Service) SummaryFile(content string) (string, error) {
	openAIClient := openai.NewClient()
	summary, err := openAIClient.Summary(content)
	if err != nil {
		return "", domain.ErrSummaryBook
	}

	return summary, nil
}
