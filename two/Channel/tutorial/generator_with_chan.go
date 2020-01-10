package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{
		"The yellow fish swins slowly",
		"The brown dog barks loudly",
		"The dark bird land on the small tree",
	}

	histogram := make(map[string]int)
	wordsCh := make(chan string)

	go func() {
		defer close(wordsCh)

		for _, line := range data {
			words := strings.Split(line, " ")

			for _, word := range words {
				word = strings.ToLower(word)
				wordsCh <- word
			}
		}
	}()

	for {
		word, opened := <-wordsCh
		if !opened {
			break
		}
		histogram[word]++
	}

	for k, v := range histogram {
		fmt.Println(fmt.Sprintf("%s\t(%d)", k, v))
	}
}
