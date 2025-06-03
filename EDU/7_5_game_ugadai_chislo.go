package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

var ErrUserEndGame = errors.New("user end game")

func generateRandomNumber(min, max int) int {
	return rand.IntN(max - min + 1) + min

}

func getUserInput(min, max int) (int, error) {
	var input string
	fmt.Printf("Ваше предположение (либо для завершения, введите слово 'выход': )")
	fmt.Scanln(&input)

	if strings.ToLower(input) == "выход" {
		return 0, ErrUserEndGame
	}
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %w", &err)
	}

	return number, nil
}

func playGame() error {
	min := 1
	max := 100
	randomNumber := generateRandomNumber(min, max)
	attempts := 0

	fmt.Println("Компьютер загадал случайное число от 1 до 100 включительно. Угадайте его!")

	for {
		input, err := getUserInput()
		if err != nil {
			if err == ErrUserEndGame {
				return err
			}
			fmt.Println("Некорректный ввод. Пожалуйста введите целое число")
			continue

		}
		attempts++

		if input < randomNumber {
			fmt.Println("Загаданное число больше")
		} else if input > randomNumber {
			fmt.Println("Загаданное число меньше")
		} else {
			fmt.Printf("Правильно! Вы угадали число с %d попытки!\n", attempts)
			break
		}
	}
	return nil

}

func printEndGameMessage() {
	fmt.Println("спасибо за игру! До свидания!")
}

func main() {
	for {
		if err := playGame(); err != nil {
			if err == ErrUserEndGame {
				printEndGameMessage()
				break
			}
			fmt.Printf("Возникла ошибка во время игры: %s\n", err)
			os.Exit(1)
		}

		var playAgain string 
			fmt.Printf("Хотите сыграть еще раз? (если хотите напишите да): ")
			fmt.Scanln(&playAgain)
			if playAgain != "да" {
				printEndGameMessage()
				break
			}
		}
}
