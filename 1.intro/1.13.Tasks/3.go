/*
Дано трехзначное число. Переверните его, а затем выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных
Выведите перевернутое число.
*/
package main

import "fmt"

func main() {
	// var x int
	// fmt.Scan(&x)
	// a, b, c := x/100, x/10%10, x%10
	// fmt.Printf("%v%v%v", c, b, a)
	stringMain()
}

func stringMain() {
	// is it possible to make something like this in python?
	// mystring[::-1]
	var x string
	var result string
	fmt.Scan(&x)
	for _, char := range x {
		result = string(char) + result
	}
	fmt.Println(result)
}
