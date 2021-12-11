package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Введите имя: ")
	// & перед переменной это ссылка на нее
	fmt.Scan(&name)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	fmt.Printf("Привет %v, вам сегодня %v лет\n", name, age)
}
