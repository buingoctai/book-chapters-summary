package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/openai/openai-go/openai"
)

func main() {
	// Replace with your OpenAI API key
	openai.SetApiKey("<YOUR_OPENAI_API_KEY>")

	// Replace with the path to your book text file
	bookFilePath := "<BOOK_FILE_PATH>"

	// Read book content
	bookText, err := ioutil.ReadFile(bookFilePath)
	if err != nil {
		fmt.Printf("Error reading book file: %v\n", err)
		return
	}

	// split total content into chapter parts
	// continue split chapter content into chunks
	// Prepare input
	summaryPrompt := "Summarize the following book content:\n" + string(bookText)

	// OpenAI request
	request := openai.CompletionRequest{
		Engine:  "text-davinci-003",
		Prompt:  summaryPrompt,
		MaxTokens: 150, // Adjust this value to control summary length
		N:        1,
		Stop:     "<|endoftext|>",
		Temperature: 0,
	}

	// Send request
	response, err := openai.Completions(request)
	if err != nil {
		fmt.Printf("Error calling OpenAI API: %v\n", err)
		return
	}

	// Process response
	if len(response.Choices) > 0 {
		summary := response.Choices[0].Text
		fmt.Println("Summary:")
		fmt.Println(summary)
	} else {
		fmt.Println("No summary generated.")
	}
}
