package gomarkov

import (
	"regexp"
	"strings"
)

// Pair is a pair of consecutive states in a sequence
type Pair struct {
	CurrentState NGram  // n = order of the chain
	NextState    string // n = 1
}

// NGram is an array of words
type NGram []string

type sparseArray map[int]int

var regNonWordChars = regexp.MustCompile("[^\\w\\s]+")
var regSeveralSpaces = regexp.MustCompile("\\s+")

func normalizeString(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)

	// Remove non-alphanumeric characters
	s = regNonWordChars.ReplaceAllString(s, "")

	// Replace consecutive spaces with a single space
	s = regSeveralSpaces.ReplaceAllString(s, " ")

	// Trim whitespace
	s = strings.TrimSpace(s)

	return s
}

func (ngram NGram) key() string {
	return normalizeString(strings.Join(ngram, " "))
}

func (s sparseArray) sum() int {
	sum := 0
	for _, count := range s {
		sum += count
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func array(value string, count int) []string {
	arr := make([]string, count)
	for i := range arr {
		arr[i] = value
	}
	return arr
}

// MakePairs generates n-gram pairs of consecutive states in a sequence
func MakePairs(tokens []string, order int) []Pair {
	var pairs []Pair
	for i := 0; i < len(tokens)-order; i++ {
		pair := Pair{
			CurrentState: tokens[i : i+order],
			NextState:    tokens[i+order],
		}
		pairs = append(pairs, pair)
	}
	return pairs
}
