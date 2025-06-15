/*
Необходимо реализовать функцию magicSort, которая принимает двумерный слайс целых чисел и 
сортирует его по следующим критериям:

Внешний слайс должен быть отсортирован по возрастанию суммы чисел в каждом внутреннем слайсе,
 если слайсы имеют одинаковую сумму, то они должны быть в том же порядке, что и были.
Каждый внутренний слайс должен быть отсортирован в порядке убывания, 
при этом четные числа должны идти перед нечетными. Ноль идет перед всеми числами.
*/

import (
	"fmt"
	"sort"
)

func magicSort(arr [][]int) {
	// Сортировка каждого внутреннего слайса
	for _, inner := range arr {
		sort.Slice(inner, func(i, j int) bool {
			a, b := inner[i], inner[j]
			// Ноль всегда первый
			if a == 0 {
				return true
			}
			if b == 0 {
				return false
			}
			// Четные перед нечетными
			aEven := a%2 == 0
			bEven := b%2 == 0

			if aEven && !bEven {
				return true
			}
			if !aEven && bEven {
				return false
			}
			return a > b	// Для чисел одной четности сортируем по убыванию
		})
	}
	// Сортировка внешнего слайса по сумме (с сохранением порядка при равных суммах)
	sort.SliceStable(arr, func(i, j int) bool{
		sumI, sumJ := 0, 0
		for _, value := range arr[i] {
			sumI += value
		}
		for _,value := range arr[j] {
			sumJ += value
		}
		return sumI < sumJ
	})

}





