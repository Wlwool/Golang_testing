/*
Вам необходимо реализовать две функции, первая - invertMap, которая будет инвертировать
входную map, вторая - printMap, которая будет выводить данные, полученные от invertMap в 
правильном формате.

Функция invertMap должна принимать map типа map[string]int и возвращать новую map 
типа map[int][]string. В новой map ключами будут значения из исходной map, 
а значениями — срезы ключей, которые соответствуют этим значениям.

Функция printMap должна принимать то, что возвращает функция invertMap, то есть, 
функция printMap должна принимать map[int][]string и выводить эти данные в 
определенном формате.
*/

import (
	"fmt"
	"sort"
)

func main () {
	m := map[string]int{
    "banana":     2,
    "apple":      1,
    "grapefruit": 3,
    "cherry":     1,
  	}
  	invertedMap := invertMap(m)
  	printMap(invertedMap)

}

func invertMap(m map[string]int) map[int][]string {
	inverted := make(map[int][]string) 
	for k, v := range m {
		inverted[v] = append(inverted[v], k)
	}
	return inverted
}

func printMap(m map[int][]string) {
	// Собираем ключи для сортировки
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	
	fmt.Println("{")
	for _, k := range keys {
		// Сортируем строки для каждого ключа
		values := m[k]
		sort.Strings(values)
		
		// Форматируем вывод
		fmt.Printf("  %d: [", k)
		for i, val := range values {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("\"%s\"", val)
		}
		fmt.Println("],")
	}
	fmt.Println("}")
}
