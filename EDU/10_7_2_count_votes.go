/*
Функция countVotes принимает один аргумент: votes типа map[string]int, 
где ключами являются имена кандидатов (типа string), а значениями — количество голосов, 
полученных каждым кандидатом (типа int). 

Функция countVotes должна возвращать имя кандидата (типа string), 
который получил наибольшее количество голосов. Если в карте нет голосов, 
функция должна возвращать пустую строку.

Примечание
Если несколько кандидатов имеют одинаковое максимальное количество голосов, 
верните имена всех кандидатов через запятую, в алфавитном порядке.
Обработайте случай, когда карта может быть пустой, 
либо когда никто не проголосовал за кандидатов.
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	m := map[string]int{
		"Linus Torvalds":  5,
		"James Gosling":   3,
		"Tim Berners-Lee": 4,
	}
	countVotes(m)

	m2 := map[string]int{
		"Mitchel Resnick": 5,  
  		"Linus Torvalds": 5,
  		"Donald Knuth": 3,
  		"Tim Berners-Lee": 5,
  		"Bjarne Stroustrup": 5,
	}
	countVotes(m2)

}

func countVotes(votes map[string]int) string {
	if len(votes) == 0 {
		return ""
	}

	maxValue := 0
	for _, rating := range votes {
		if rating > maxValue {
			maxValue = rating
		}
	}

	var topNames []string
	for name, value := range votes {
		if value == 0 {
			return ""
		}
		if value == maxValue {
			topNames = append(topNames, name)
		}
	}
	sort.Strings(topNames)
	return strings.Join(topNames, ", ")
}
