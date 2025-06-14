import (
	"errors"
	"fmt"
)



func main () {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4, 5, 6, 7}
	result, err := intersectSlices(slice1, slice2)
	fmt.Println("Test 1:", result, "Error:", err) 

	// Тест 2: Нет пересечений
	slice3 := []int{1, 3, 5, 7}
	slice4 := []int{2, 4, 6, 8}
	result, err = intersectSlices(slice3, slice4)
	fmt.Println("Test 2:", result, "Error:", err) 

	// Тест 3: Один слайс = nil
	var nilSlice []int = nil
	slice5 := []int{1, 2, 3}
	result, err = intersectSlices(nilSlice, slice5)
	fmt.Println("Test 3:", result, "Error:", err) 

	// Тест 4: Оба пустые
	empty1 := []int{}
	empty2 := []int{}
	result, err = intersectSlices(empty1, empty2)
	fmt.Println("Test 4:", result, "Error:", err)

	// Тест 5: Отрицательные числа
	slice6 := []int{-3, -2, -1, 0, 1, 2, 3}
	slice7 := []int{-2, 0, 2, 4, 6}
	result, err = intersectSlices(slice6, slice7)
	fmt.Println("Test 5:", result, "Error:", err) 

	// Тест 6: Неполное совпадение
	slice8 := []int{1, 2, 3, 5}
	slice9 := []int{3, 4, 5, 6, 7}
	result, err = intersectSlices(slice8, slice9)
	fmt.Println("Test 6:", result, "Error:", err) 
}


func intersectSlices(s1, s2 []int) ([]int, error) {
	if s1 == nil || s2 == nil {
		return nil, errors.New("slices cannot be nil")
	}

	var result []int 	// Создаем слайс для результата
	i, j := 0, 0		// Два указателя для прохода по слайсам

	// Пока не достигли конца одного из слайсов
	for i < len(s1) && j < len(s2) {
		switch {
		case s1[i] == s2[j]:
			result = append(result, s1[i])		// Элементы совпадают - добавляем в результат
			i++
			j++
		case s1[i] < s2[j]:
			i++			// Элемент в первом слайсе меньше - двигаем первый указатель
		default:
			j++			// Элемент во втором слайсе меньше - двигаем второй указатель
		}
	}
	return result, nil
}
