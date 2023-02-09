package nlp

import (
	"regexp"
	"strings"
)

var (
	wordRe = regexp.MustCompile(`[a-zA-Z]+`)
)


func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	var tokens []string
	for _, w := range words {
		token := strings.ToLower(w)
		tokens = append(tokens, token)
	}

	return tokens
}

/*

	text := "What's on second?"
	expected := []string{"what", "on", "second"}
	tokens := Tokenize(text)

	if tokens != expected { // 在 Go 中，不能用 == 比较 slice（nil 除外）

	*/ 