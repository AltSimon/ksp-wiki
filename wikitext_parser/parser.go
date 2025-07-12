package wikitext_parser

import (
	"log"
	"strings"
)

// Parses wikitext to markdown
func Parse(lines []string) []string {
	result := []string{}
	for _, line := range lines {
		first_character := line[0]
		switch first_character {
		case '=':
			result = append(result, HeadingToMarkdown(line))
		}
	}
	return result
}

// Checks if the heading provided is using symmetrical prefix and suffix of size provided
func isCorrectHeadingSuffixSize(heading string, size int) bool {
	isCorrectPrefix := strings.HasPrefix(heading, strings.Repeat("=", size)) && !strings.HasPrefix(heading, strings.Repeat("=", size+1))
	isCorrectSuffix := strings.HasSuffix(heading, strings.Repeat("=", size)) && !strings.HasSuffix(heading, strings.Repeat("=", size+1))
	return isCorrectPrefix && isCorrectSuffix
}

// Gets the size of the heading provided, e.g. == Heading == will be 2
func getHeadingNumber(heading string) int {
	if !strings.HasPrefix(heading, "=") {
		log.Println("not a heading")
		return -1
	}
	size := 0
	for index, char := range heading {
		if char != '=' {
			size = index
			break
		}
	}
	if !isCorrectHeadingSuffixSize(heading, size) {
		log.Println("heading suffix or prefix is not correct")
		return -1
	}
	if size >= 6 {
		log.Println("heading size is too small, using 6")
		return 6
	}
	return size
}

// Transforms a wikitext heading to a markdown heading
func HeadingToMarkdown(heading string) string {
	size := getHeadingNumber(heading)
	prefix := strings.Repeat("=", size)
	heading = strings.TrimPrefix(heading, prefix)
	heading = strings.TrimSuffix(heading, prefix)
	heading = strings.TrimSpace(heading)
	line := ""
	if size <= 2 {
		line = "\n---"
	}
	return strings.Repeat("#", size) + " " + heading + line
}
