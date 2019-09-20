package analize

import (
	"sort"
	"strings"
	"unicode"
)

const countTop = 10

// MostCommonWords finds the most common words in text
func MostCommonWords(text string) []string {
	if len(text) == 0 {
		return nil
	}

	words := strings.FieldsFunc(text, func(c rune) bool {
		return !unicode.IsLetter(c)
	})

	vocabulary := make(map[string]int)
	for _, word := range words {
		vocabulary[strings.ToLower(word)]++
	}

	topWords := make([]string, 0, len(vocabulary))
	for key := range vocabulary {
		topWords = append(topWords, key)
	}

	if len(topWords) <= countTop {
		return topWords
	}

	sort.Slice(topWords, func(i, j int) bool {
		return vocabulary[topWords[i]] > vocabulary[topWords[j]]
	})

	return topWords[:countTop]
}
