/*
Доступные товары и их цены:
"Клавиатура JZ9": 19200
"Наушники N45": 9600
"Смартфон S10": 55000
Если пользователь введет название товара, которого нет в списке, программа должна вывести сообщение:

Товар "название_введенного_товара" не найден.
Поиск должен быть нечувствителен к регистру. Например, если пользователь введет "наУШНИКИ n45", программа должна вернуть цену 9600.

Примеры:
Введите название товара: наУШНИКИ n45
Наушники N45: 9600
Введите название товара: сМАРТ
Смартфон S10: 55000
Введите название товара: планшет
Товар "планшет" не найден.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)


func goods() {
	fmt.Printf("Введите название товара: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: %s", err)
	}

	inputLower := strings.ToLower(scanner.Text())
	
	// if strings.Contains("клавиатура JZ9", inputLower) {
	// 	fmt.Println("Клавиатура JZ9: 19200")
	// } else if strings.Contains("наушники N45", inputLower) {
	// 	fmt.Println("Наушники N45: 9600")
	// } else if strings.Contains("смартфон S10", inputLower) {
	// 	fmt.Println("Смартфон S10: 55000") 
	// }else {
	// 	fmt.Printf("Товар %q не найден.\n", inputLower)
	// }

	switch {
	case strings.Contains("клавиатура JZ9", inputLower):
		fmt.Println("Клавиатура JZ9: 19200")
	case strings.Contains("наушники N45", inputLower):
		fmt.Println("Наушники N45: 9600")
	case strings.Contains("смартфон S10", inputLower):
		fmt.Println("Смартфон S10: 55000")
	default: 
		fmt.Printf("Товар %q не найден.\n", inputLower)
	}

}
