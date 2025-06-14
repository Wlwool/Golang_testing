func main () {
	input := []int {1, 2, 6, 11, 8}
	result := PlayWithSlice(input)
	fmt.Println("Result: ", result)
}

func PlayWithSlice(s []int) []int {
	copySlice := make([]int, len(s))  // клон слайса
	copy(copySlice, s)

	// вставка числа 10, поиск с конца есои первое число > 10
	insertIndex := -1
	for i := len(copySlice) - 1; i>=0; i-- {
		if copySlice[i] > 10 {
			insertIndex = i
			break
		} 
	}

	// Вставляем 100 после найденного элемента
	if insertIndex != -1 {
		copySlice = append(copySlice[:insertIndex + 1], append([]int{100}, copySlice[insertIndex+1:]...)...)
	}

	// eсли сумма всех чисел больше 100, добавляется число 500 в конец слайса
	sum :=0
	for _, num := range copySlice {
		sum += num
	}
	if sum > 100 {
		copySlice = append(copySlice, 500)
	}
	
	// tсли в оригинальном слайсе четных чисел больше, чем нечетных, 
	// вставляется число 1000 в начало клона слайса.
	evenCount, oddCount := 0, 0
	for _, num := range s {
		if num % 2 == 0{
			evenCount++
		} else {
			oddCount++
		}
	}

	if evenCount > oddCount {
		copySlice = append([]int{1000}, copySlice...)
	}

	return copySlice
}