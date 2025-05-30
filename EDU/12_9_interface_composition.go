package main

import (

	"fmt"

)


type Writer interface {
	Write(string)
}

type DB interface {
	Writer
	Read() string
}

type MyDB struct {}
func (m MyDB) Write(str string) {
	fmt.Printf("Записали %q", str)
}

func (m MyDB) Read() string {
	return "Возвращаем нужную строку"
}

type Logger struct {}
func (m Logger) Write(str string) {
	fmt.Printf("Залогировали %q\n", str)
}

// func main() {
// 	myDB := MyDB{}
// 	logger := Logger {}
// 	run(myDB, logger)
// }

func run (db DB, logger Writer) {
	db.Write("Привет")
	fmt.Println(db.Read())

	logger.Write("Логи")
}
