package fake

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var wordList = []string{
	"apple",
	"banana",
	"orange",
	"dog",
	"cat",
	"horse",
	"cow",
}

func randomWord(prefix string) string {
	return fmt.Sprintf("%s%d", prefix, time.Now().UnixNano())
}

// TODO: I'd like this to make things more sentence looking (e.g. capitalized
// first letter, period at the end, etc.) but it doesn't really matter.
func randomSentence(length int) string {
	words := []string{}
	for i := 0; i < length; i++ {
		word := wordList[rand.Intn(len(wordList))]
		words = append(words, word)
	}
	return strings.Join(words, " ")
}
