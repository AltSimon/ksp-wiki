package wikitext_parser

import (
	"log"
	"strings"
)

func isCorrectHeadingSuffixSize(heading string, size int) bool {
	return strings.HasSuffix(heading, strings.Repeat("=", size)) && !strings.HasSuffix(heading, strings.Repeat("=", size+1))
}

func getHeadingNumber(heading string) int {
	if !strings.HasPrefix(heading, "=") {
		log.Fatal("not a heading")
	}
	size := 0
	for index, char := range heading {
		if char != '=' {
			size = index
			break
		}
	}
	if !isCorrectHeadingSuffixSize(heading, size) {
		log.Fatal("heading suffix is not the same as the prefix")
	}
	if size >= 6 {
		log.Println("heading size is too small, using 6")
		return 6
	}
	return size
}

func HeadingToMarkdown(heading string) string {
	size := getHeadingNumber(heading)
	prefix := strings.Repeat("=", size)
	heading = strings.TrimPrefix(heading, prefix)
	heading = strings.TrimSuffix(heading, prefix)
	line := ""
	if size <= 2 {
		line = "\n---"
	}
	return strings.Repeat("#", size) + heading + line
}

// TODO: Tests !!!
