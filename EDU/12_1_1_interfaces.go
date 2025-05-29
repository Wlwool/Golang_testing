package main

import (
	"fmt"
)


type Dog struct {}
func (Dog) Speak() string {
	return "Waf!!!"
}
func (Dog) Growl() string{
	return "RrRrRrRr!"
}

type Cat struct {}
func (Cat) Speak() string {
	return "Weow!"
}

type Bird struct {}
func (Bird) Speak() string {
	return "chirickkk!"
}

// func main() {
// 	dog := Dog{}
// 	cat := Cat{}
// 	bird := Bird{}
	
// 	makeSound(dog)
// 	makeSound(cat)
// 	makeSound(bird)
// }

type Speaker interface {
	Speak() string
}

func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}
