package main

import (
	"fmt"
	"maps"
)

/*
Дано два map, где ключи предвствляют собой строки, а значения - срехзы  целых чисел ([]int).
Необходимо сравнить эти два map по следующим критериям:

1. Сравнение осуществляется  по ключам. Если ключ присутствуеи в одном, но отсутствует в другом,
то map не равны.
2. Для каждого ключа, который присутствует в обоих map, необходимо сравнить суммы значений
соответствующих срезов. Если суммы равны, то срезы считаются равными для данного ключа.
*/
func CompareMaps() {
	m1 := map[string][]int{
		"a": {5, 2},
		"b": {20, 11},
	}
	m2 := map[string][]int{
		"a": {1, 5, 1},
		"b": {11, 20},
	}

	result := maps.EqualFunc(m1, m2, func(v1[]int, v2 []int) bool {
		sum1 := 0
		for _, v := range v1 {
			sum1 += v
		}
		sum2 := 0
		for _, v := range v2 {
			sum2 += v
		}
		return sum1 == sum2
	})
	fmt.Println(result)
}