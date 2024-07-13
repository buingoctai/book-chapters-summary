package templates

import (
	"strings"
)

// import "regexp"

type Template interface {
    AdaptContent(content string) []string
}


type WordBasedTemplate struct{}

func (t *WordBasedTemplate) AdaptContent(content string) []string {
    // Implement adaptation logic for word-based separators


    chapters := make(map[string]string)

	lines := strings.Split(content, "\n")
	var currentChapter string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.HasPrefix(line, " ") {
			currentChapter += line + "\n"
		} else {
			if currentChapter != "" {
				chapters[strings.TrimSpace(strings.Split(currentChapter, "\n")[0])] = currentChapter
			}
			currentChapter = line + "\n"
		}
	}

	// Add the last chapter
	if currentChapter != "" {
		chapters[strings.TrimSpace(strings.Split(currentChapter, "\n")[0])] = currentChapter
	}

	chaptersList := []string{} // create a list to store chapter content

	for _, chapterContent := range chapters {
		chaptersList = append(chaptersList, chapterContent) // append chapter content to the list
	}

	return chaptersList // return the list of chapter content
}

type LineBasedTemplate struct{}

func (t *LineBasedTemplate) AdaptContent(content string) []string {
    // Implement adaptation logic for line-based separators
    return []string{}
}
