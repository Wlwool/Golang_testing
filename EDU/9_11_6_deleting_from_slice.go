/*
Необходимо создать функцию с названием DeletingFromSlice, 
которая принимает слайс целых чисел и возвращает новый слайс целых чисел. 
Функция должна выполнять следующие действия:
- Удалить из принятого слайса последнее значение, если это значение существует и оно больше 10.
- Удалить значение по индексу 2, если такое значение есть и вместимость слайса больше 5.
- Удалить первое значение из слайса, если оно присутствует и были удалены значения,
  указанные в первых двух пунктах.
- Убрать лишнюю вместимость у слайса.
Функция должна вернуть полученный слайс.
*/

func main () {
	input := []int {1, 2, 3, 4, 5, 11} 
	result := DeletingFromSlice(input)
	fmt.Println("Result: ", result)

	input2 := []int {1, 2, 3, 4, 5, 6} 
	result2 := DeletingFromSlice(input2)
	fmt.Println("Result: ", result2)
}

func DeletingFromSlice(s []int) []int {
	copySlice := make([]int, len(s))
	copy(copySlice, s)

	deletedLAst := false  // флаг для отслеживания удаление последнего элемента
	deletedIndex := false  // отслеживает удаление по индексу 2

	// если значение существует и больше 10 удаление последнего значения
	if len(copySlice) > 0 && copySlice[len(copySlice)-1] > 10 {
		copySlice = copySlice[:len(copySlice)-1]
		deletedLAst = true
	}
	
	// удаление значения по индексу 2, если вместимость > 5 и  индекс существует
	if cap(copySlice) > 5 && len(copySlice) > 2 {
		copySlice = append(copySlice[:2], copySlice[3:]...)
		deletedIndex = true
	}

	if deletedLAst && deletedIndex && len(copySlice) > 0 {
		copySlice = copySlice[1:]
	}

	result := make([]int, len(copySlice))  // убираем лишнюю совместимость
	copy(result, copySlice)

	return result
}



