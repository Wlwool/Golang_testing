package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs()) // Output: 50
	// fmt.Println(f.Abs()) // Output: 5

}

















func CompareMaxValues(map1, map2 map[string][]int) bool {
	for key, slice1 := range map1 {
		slice2, exists := map2[key]
		if !exists{
			return false
		}

		if len(slice1) == 0 && len(slice2) == 0 {
			continue
		}
		if len(slice1) == 0 || len(slice2) == 0 {
			return false
		}

		max1 := slice1[0]
		for _, val := range slice1{
			if val > max1 {
				max1 = val
			}
		}

		max2 := slice2[0]
		for _, val := range slice2{
			if val > max2 {
				max2 = val
			}
		}

		if max1 != max2 {
			return false
		}
	}
	for key := range map2 {
		if _, exists := map1[key]; !exists {
			return false
		}
	}

	return true
}



func CountMaxFrequency(nums []int) int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}

	maxFrequency := 0
	for _, freq := range count {
		if freq > maxFrequency {
			maxFrequency = freq
		}
	}
	return maxFrequency

}


