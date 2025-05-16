package main

import (
	"fmt"
)

// func main() {
// 	arr1 := [10]int{1, 2, 3, 4, 5, 4, 3, 2, 1, 1}
//     isPalindrome(arr1) 

//     arr2 := [10]int{1, 2, 3, 2, 1, 1, 2, 3, 2, 1}
//     isPalindrome(arr2) 

//     arr3 := [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
//     isPalindrome(arr3) 

//     arr4 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//     isPalindrome(arr4)
// }

func isPalindrome(numbers [10]int) {
	for i := 0; i < 5; i++ {
		if numbers[i] != numbers[9-i] {
			fmt.Println("Не палиндром!")
			return
		}
	}
	fmt.Println("Это палиндром!")
}

// func isPalindrome(numbers [10]int) {
// 	for i := 0; i < 5; i++ {
// 		v := 9 - i
// 		if numbers[i] != numbers[v] {
// 			fmt.Println("Не палиндром!")
// 			return
// 		}
// 	}
// 	fmt.Println("Это палиндром!")
// }
