package rest

import (
	"log"
	"mime/multipart"
	"strings"
	"sync"
	"time"

	"github.com/buingoctai/book-chapters-summary/domain"
	"github.com/buingoctai/book-chapters-summary/internal/rest/validator"
	"github.com/buingoctai/book-chapters-summary/pkg/response"
	"github.com/buingoctai/book-chapters-summary/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	openAICallingConcurrency = 5
	numberOfRequestPerSecond = 10
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

type BookService interface {
	ReadUploadedFile(file multipart.File) ([]byte, error)
	UploadFile(file_bytes []byte, fileName string) (string, error)
	LoadFile(name string) (string, error)
	SummaryFile(content string) (string, error)
}

type BookHandler struct {
	Service BookService
}

func NewBookHandler(e *echo.Echo, svc BookService) {
	handler := &BookHandler{
		Service: svc,
	}
	e.POST("/book/upload", handler.UploadBook)
	e.GET("/book/summary", handler.GetBookSummary)
}

func (a *BookHandler) UploadBook(c echo.Context) error {
	formFile, err := c.FormFile("file")
	if err != nil {
		return c.JSON(response.GetStatusCode(domain.ErrNotFoundFile), response.GetError(domain.ErrNotFoundFile))
	}

	file, err := formFile.Open()
	if err != nil {
		return c.JSON(response.GetStatusCode(domain.ErrBadParamFile), response.GetError(domain.ErrBadParamFile))
	}

	defer file.Close()

	isValid, err := validator.IsBookValid(formFile)
	if !isValid {
		return c.JSON(response.GetStatusCode(err), response.GetError(err))
	}
	

	fileBytes, err := a.Service.ReadUploadedFile(file)
	if err != nil {
		return c.JSON(response.GetStatusCode(err), response.GetError(err))
	}

	localPath, err := a.Service.UploadFile(fileBytes, formFile.Filename)

	book := domain.Book{
		Id:     0,
		Name:   formFile.Filename,
		Owner:  "",
		Url:    localPath,
		SummarizedUrl: "",
	}

	if err != nil {
		return c.JSON(response.GetStatusCode(err), response.GetError(err))
	}

	return c.JSON(response.GetStatusCode(nil), book)
}

func (a *BookHandler) GetBookSummary(c echo.Context) error {
	// name := c.QueryParam("name");

	// if name == "" {
	// 	return c.JSON(http.StatusBadRequest, ResponseError{Message: "Filename is required"})
	// }

	name := "test-copy.txt"

	content, err := a.Service.LoadFile(name)
	if err != nil {
		return c.JSON(response.GetStatusCode(err), response.GetError(err))
	}

	chapters := utils.SplitIntoChapters(content, "word")
	logrus.Info(chapters, len(chapters))	

    semaphore := make(chan struct{}, openAICallingConcurrency)
    var wg sync.WaitGroup
    var summaries []string
    var mu sync.Mutex

	// Define rate limit (e.g., 10 requests per second)
	rateLimit := time.Second / numberOfRequestPerSecond
    ticker := time.NewTicker(rateLimit)
    defer ticker.Stop()
	
    for _, chapter := range chapters {
        wg.Add(1)
        go func(chapter string) {
            defer wg.Done()
			<-ticker.C // Wait for the ticker
            semaphore <- struct{}{}
            defer func() { <-semaphore }()

            summary, err := a.Service.SummaryFile(chapter)
            if err != nil {
                log.Printf("Failed to summarize chapter: %v", err)
                return
            }

            mu.Lock()
            summaries = append(summaries, summary)
            mu.Unlock()
        }(chapter)
    }

    wg.Wait()


	summaryBookBytes := []byte(strings.Join(summaries, "\n"))
	summarizedName := "summarized-" + name
	localPath, err := a.Service.UploadFile(summaryBookBytes, summarizedName)

	book := domain.Book{
		Id:     0,
		Name:   name,
		Owner:  "",
		Url:    "",
		SummarizedUrl: localPath,
	}

	if err != nil {
		return c.JSON(response.GetStatusCode(err), response.GetError(err))
	}

	return c.JSON(response.GetStatusCode(nil), book)
}
