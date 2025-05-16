package main

import (
	"fmt"
	// "os"
)


// func main() {
// 	score, err := GetScore()
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		os.Exit(1)
// 	}
// 	grade := LetterGrade(score)
// 	fmt.Printf("Введите оценку (0 - 100): %s\n", grade)
// }

func GetScore() (int, error) {
	var score int
	fmt.Println("введите количестов баллов (0-100): ")

	_, err := fmt.Scanln(&score)
	if err != nil{
		return 0, fmt.Errorf("некорректный ввод: %w", err)
	}

	if score < 0 || score > 100{
		return 0, fmt.Errorf("оценка должна быть в диапазоне от 0 до 100")
	}

	return score, nil
}

func LetterGrade(score int) string {
	switch{
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}
