package main

import (
	"fmt"
	"math/rand/v2"
)



func main() {
	num := 8
	rollDice(num)

}

func rollDice(num int) {
	attempt := 0  // счетчик попыток бросков

	for {
		// генерация рандомных значений для кубиков в пределах (1-6)
		die1 := rand.IntN(6)+1 
		die2 := rand.IntN(6)+1 
		sum := die1 + die2 
		attempt++

		// базовое сообщение
		message := fmt.Sprintf("Выпало %d и %d, в сумме %d", die1, die2, sum)

		if sum == num {
			var attempsWord string
			if attempt%10 == 1 && attempt%100 != 11 {
				attempsWord = "потребовался"
			} else {attempsWord = "потребовалось"}
			
			// определение правильного окончания для слова "бросок"
			var nounWord string
			switch {
			case attempt%10 == 1 && attempt%100 != 11:
				nounWord = "бросок"
			case attempt%10 >= 2 && attempt%10 <= 4 && (attempt%100 < 10 || attempt%100 >= 20):
				nounWord = "броска"
			default:
				nounWord = "бросков"
			}

			fmt.Printf("%s, на это %s %d %s.\n", message, attempsWord, attempt, nounWord)
			break
		} else {
		fmt.Printf("%s, бросаем еще раз.\n", message)
		}
	}
}