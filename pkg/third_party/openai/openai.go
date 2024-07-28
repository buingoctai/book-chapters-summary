package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/buingoctai/book-chapters-summary/domain"
)

const (
    maxRetries   = 5
)

type OpenAI interface {
	Summary(content string) (string, error)
}

type MessageItem struct {
	Role   string `json:"role"`
	Content string `json:"content"`
}

type Client struct{}

func NewClient() OpenAI {
	return &Client{}
}

func (c *Client) Summary(content string) (string, error) {
	data := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []MessageItem{
			{
				Role:   "system",
				Content: "Summarize content you are provided with for a second-grade student.",
			},
			{
				Role:   "user",
				Content: content,
			},
		},
		"temperature": 0.7,
		"max_tokens": 64,
		"top_p": 1,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", os.Getenv("OPEN_AI_ENDPOINT"), bytes.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPEN_AI_KEY")))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	var resp *http.Response
	for i := 0; i < maxRetries; i++ {
        resp, err = client.Do(req)
        if err != nil {
            return "", err
        }
        defer resp.Body.Close()

        if resp.StatusCode == http.StatusOK {
            break
        } else if resp.StatusCode == http.StatusTooManyRequests {
            log.Println("Too Many Requests. Retrying...")
            time.Sleep(time.Duration(2<<i) * time.Second) // Exponential backoff
        } else {

            return "", fmt.Errorf("unexpected status code: %v", resp.StatusCode)
        }
    }


	if resp.StatusCode != http.StatusOK {
		return "", domain.ErrOpenAIService
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return "", err
	}

	if _, ok := response["choices"]; !ok {
		return "", domain.ErrOpenAIService
	}

	choices := response["choices"].([]interface{})
	if len(choices) == 0 {
		return "", domain.ErrOpenAIService
	}

	choice := choices[0].(map[string]interface{})

	summary := choice["message"].(map[string]interface{})["content"].(string)

	// Clean summary (remove extra characters)
	summary = strings.TrimSpace(summary)

	return summary, nil
}

