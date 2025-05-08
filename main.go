package main

import (
	"fmt"
	"strings"
)


func main() {
	products := map[string]int{
		"Клавиатура JZ9": 19200,
		"Наушники N45":   9600,
		"Смартфон S10":   55000,
	}
	fmt.Println("Введите название товара: ")
	var tovar string
	fmt.Scanln(&tovar)

	inputLower := strings.ToLower(tovar)

	found := false

	for name, price := range products{
		if strings.ToLower(name) == inputLower{
			fmt.Printf("%s: %d\n", name, price)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Товар \"%s\" не найден.\n", tovar)
	}

}

