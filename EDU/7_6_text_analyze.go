/*
Анализ строки
Ваша задача — разработать программу, которая будет анализировать текст, введенный пользователем. 
Программа должна выполнять несколько ключевых функций, позволяющих получить различные 
статистические данные о тексте.

Ввод текста
Создайте функцию GetInput() (string, error), которая запрашивает у пользователя ввод текста и 
возвращает его для дальнейшего анализа. Если функция вернет ошибку, программу необходимо завершить.

Подсчет символов
Реализуйте функцию CountCharacters(text string) (letters, digits, spaces, punctuation int), 
которая принимает текст и возвращает количество:

букв (letters)
цифр (digits)
пробелов (spaces)
знаков препинания (punctuation)

Вывод результатов
Напишите функцию DisplayResults(letters, digits, spaces, punctuation int), 
которая принимает результаты анализа и выводит их на экран в удобочитаемом формате. Например:

Количество букв: 50
Количество цифр: 10
Количество пробелов: 5
Количество знаков препинания: 3
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)



func GetInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Введите текст: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("некорректный ввод %w", err)
	}
	return input, nil
}

func CountCharacters(text string) (letters, digits, spaces, punctuation int) {
	for _, r :=range text {
		switch {
		case unicode.IsLetter(r):
			letters++
		case unicode.IsDigit(r):
			digits++
		case r == ' ' || r == '\t':
			spaces++
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			punctuation++
		}
	}
	return 
}

func DisplayResults(letters, digits, spaces, punctuation int) {
	fmt.Println("Количество букв: ", letters)      
	fmt.Println("Количество цифр: ", digits)         
	fmt.Println("Количество пробелов: ", spaces)         
	fmt.Println("Количество знаков препинания: ", punctuation) 
}

func main () {
	input, err := GetInput()
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}

	if len(input) == 0 || input == "\n" {
		fmt.Println("Вы ввели пустую строку!")
		DisplayResults(0, 0, 0, 0)
		return
	}
	letters, digits, spaces, punctuations := CountCharacters(input)
	DisplayResults(letters, digits, spaces, punctuations)
}
