/*
Входные данные
Функция принимает массив целых чисел фиксированного размера — 5 значений, и возвращает строку, 
представляющую сгенерированный код.

Генерация секретного кода
В начале кода должно находиться минимальное значение массива.
В конце кода должно находиться максимальное значение массива.
Между максимальным и минимальными значениями, должны быть расположены все значения, 
перед которыми должна быть проставлена буква.
Если элемент четный, перед ним добавляется префикс "E" (от слова Even — четный).
Если элемент нечетный, перед ним добавляется префикс "O" (от слова Odd — нечетный).
Пример
Для входного массива [3, 8, 1, 8, 1]:

Минимальное значение: 1
Максимальное значение: 8
Префиксы: O3, E8, O1

Сгенерированный код будет: 1O3E8O1E8O18
*/

import (
	"fmt"
	"strconv"
	"strings"
)

func main () {
	input := []int {3, 8, 1, 8, 1}
	result := generateCode(input)
	fmt.Println("Code: ", result)

	input2 := []int {3, 8, 1, 8, 1}
	result2 := generateCodeVersion2(input2)
	fmt.Println("Code: ", result2)

}

func generateCode (nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	// определение минимального и максимального значений
	maxVal, minVal := nums[0], nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
		if num < minVal {
			minVal = num
		}
	}

	// Builder для эффективной конкатенации строк
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(minVal))

	// определение префикса на основе четности
	for _, num := range nums {
		prefix := "O"
		if num % 2 == 0 {
			prefix = "E"
		}
		// добавление элемента с префиксом
		builder.WriteString(prefix)
		builder.WriteString(strconv.Itoa(num))
	}
	// добавление максимального значения в конец
	builder.WriteString(strconv.Itoa(maxVal))
	
	return builder.String()
}

func generateCodeVersion2 (arr []int) string {
	    if len(arr) == 0 {
        return ""
    }
    
    minVal, maxVal := arr[0], arr[0]
    parts := make([]string, 0, len(arr)+2)
    
    for _, num := range arr {
        if num < minVal {
            minVal = num
        }
        if num > maxVal {
            maxVal = num
        }
    }
    
    parts = append(parts, strconv.Itoa(minVal))
    
    for _, num := range arr {
        prefix := "O"
        if num%2 == 0 {
            prefix = "E"
        }
        parts = append(parts, prefix+strconv.Itoa(num))
    }
    
    parts = append(parts, strconv.Itoa(maxVal))
    
    return strings.Join(parts, "")
}
