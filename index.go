package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	// "github.com/textminer/textminer/v2/sentences"
	// "github.textminer/textminer/v2/stopwords"
)

// OpenAI API details (replace with your actual API key and endpoint)
const (
	openAIKey   = "sk-team2024-home-test-vP00R3HgNtYAroFA1rRbT3BlbkFJ6XTugV2XnCDwPKOnwg18"
	openAIEndpoint = "https://api.openai.com/v1/completions"
)

// summarizeChapter takes chapter text and returns a summary using OpenAI API
func summarizeChapter(text string) (string, error) {
	// Preprocess text
	// text = strings.ToLower(text)
	// sentences, err := sentences.SentenceSegment(text, "en")
	// if err != nil {
	// 	return "", err
	// }

	// Remove stop words
	// stopWords := stopwords.LoadStopwords("en")
	// var chapterText string
	// for _, sentence := range sentences {
	// 	if !stopwords.IsStopword(sentence) {
	// 		chapterText += sentence + " "
	// 	}
	// }

	// Prepare OpenAI request data
	data := map[string]interface{}{
		"model":  "text-davinci-003", // This model is recommended by OpenAI as summarization model. It is not the same as the model above, so it should work fine.
		"prompt": "Summarize the following chapter: \n" + text,
		"max_tokens": 150, // Adjust number of tokens for desired summary length (OpenAI charges per token)
		"temperature": 0.5,
		// "n":        1,     // Request only 1 completion (the summary)
	}

	// Create JSON payload
	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Send HTTP request to OpenAI API
	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", openAIKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	fmt.Println("summarizeChapter, resp", resp, "err", err)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Parse response and extract summary
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if _, ok := response["choices"]; !ok {
		return "", fmt.Errorf("OpenAI response format error")
	}

	choices := response["choices"].([]interface{})
	if len(choices) == 0 {
		return "", fmt.Errorf("OpenAI failed to generate summary")
	}

	choice := choices[0].(map[string]interface{})
	summary := choice["text"].(string)

	// Clean summary (remove extra characters)
	summary = strings.TrimPrefix(summary, "\n")
	summary = strings.TrimSpace(summary)

	return summary, nil
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	// Read uploaded file
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error reading file: %v", err)
		return
	}
	defer file.Close()

	// Read file bytes
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error processing file: %v", err)
		return
	}

	// Save the file (replace with your logic for saving)
	err = ioutil.WriteFile("uploaded.txt", fileBytes, 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error saving file: %v", err)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully!")
}


func main() {
	// Read book content from file
	// content, err := ioutil.ReadFile("book.txt")
	// if err != nil {
	// 	fmt.Println("Error reading book file:", err)
	// 	return
	// }

	// Split book content into chapters (logic depends on chapter separators in your book file)
	// chapters := strings.Split(string(content), "\n\n") // Adjust chapter separator as needed

	// chapters := []string{
	// 	"Earth is rounded into an ellipsoid with a circumference of about 40,000 km. It is the densest planet in the Solar System. Of the four rocky planets, it is the largest and most massive. Earth is about eight light-minutes away from the Sun and orbits it, taking a year (about 365.25 days) to complete one revolution",
	// }

	// var summaries []string
	// for _, chapter := range chapters {
	// 	// Summarize each chapter
	// 	summary, err := summarizeChapter(chapter)


	// 	if err != nil {
	// 		fmt.Println("Error summarizing chapter:", err)
	// 		continue // Skip to next chapter on error
	// 	} 

	// 	fmt.Println("Result summarizing chapter:", summary)
	// 	summaries = append(summaries, summary)
	// }

	// Write summaries to a file
	// summaryContent := strings.Join(summaries, "\n\n")
	// err = ioutil.WriteFile("summaries.txt", []byte(summaryContent), 0644)


	http.HandleFunc("/upload", uploadHandler)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
