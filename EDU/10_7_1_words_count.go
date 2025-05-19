package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countWords() {
	text := "Да здравствует прекрасный язык, да здравствует golang!"
	
	strings.Map(func(r rune) rune {
		if unicode.IsPunct(r){
			return -1
		}
		return unicode.ToLower(r)
	}, text)
	
	words := strings.Split(text, " ")
	fmt.Println(words)
	
	wordsCount := make(map[string]int)
	for _, word := range words{
		wordsCount[word]++
	}

	for word, count := range wordsCount{
		fmt.Printf("Слово %q встречается в строке %d раз(а)\n", word, count)
	}
}
