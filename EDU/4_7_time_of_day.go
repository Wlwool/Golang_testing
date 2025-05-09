package main

import (
	"fmt"
	"os"
)

func main() {
	var input int
	fmt.Println("Введите время в часах (0-23): ")
	_, err := fmt.Scan(&input)

	if err != nil || input < 0 || input > 23 {
		fmt.Println("Неверно задано время")
		os.Exit(1)
	} 

	var timeOfDay string
	switch {
	case input < 6 || input > 23:
		timeOfDay = "ночь"
	case input < 12:
		timeOfDay = "утро"
	case input < 18:
		timeOfDay = "день"
	default:
		timeOfDay = "вечер"
	}

	fmt.Printf("\nСейчас: %02dч. - %s\n", input, timeOfDay)
	
}
