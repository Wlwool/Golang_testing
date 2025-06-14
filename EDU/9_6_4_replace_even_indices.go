/*
Напишите функцию replaceEvenOnEvenIndices, которая принимает двумерный слайс целых чисел и 
возвращает новый двумерный слайс целых чисел.

В возвращаемом слайсе должны быть те же значения, что у принятого слайса, 
однако, все четные числа, находящиеся на четных индексах (включая 0, в каждом из 
собственных внутренних слайсов), должны быть заменены на 0.
*/

func main () {
	input := [][]int {
		{1, 2, 3},
		{},
		{7 , 8, 9, 22, 48},
		{10 , 11},
	}
	result := replaceEvenOnEvenIndices(input)
	fmt.Println("Result: ", result)
}

func replaceEvenOnEvenIndices(v [][] int) [][]int {
	res := make([][]int, len(v))

	for i, sl := range v {
		newSlice := make([]int, len(sl))
		for j, value := range sl {
			if j%2 == 0 && value % 2 == 0 {
				newSlice[j] = 0
			} else {
				newSlice[j] = value
			}
		}
		res[i] = newSlice
	}
	return res
}
 