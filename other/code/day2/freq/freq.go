package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)


func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer file.Close()

	fmt.Println(mostCommon(file))
}

func mostCommon(r io.Reader) (string, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return "", err
	}
	return maxWord(freqs)
}

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)


func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0 {
		return "", fmt.Errorf("empty map")
	}

	// 违禁词过滤，不再统计范围内
	delete(freqs, "fuck")
	
	maxW, maxN := "", 0
	for word, count := range freqs {
		if count > maxN {
			maxW, maxN = word, count
		}
	}
	
	return fmt.Sprintf("%s - %d", maxW, maxN), nil
}


// 词频
func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int)  // word -> count
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1)  // 当前行，处理
		for _, w := range words {
			freqs[strings.ToLower(w)]++
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return freqs, nil
}