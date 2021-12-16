/*
Дано трехзначное число. Найдите сумму его цифр.

Формат входных данных
На вход дается трехзначное число.

Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var x int
	// fmt.Scan(&x)
	// a, b, c := x/100, x/10%10, x%10
	// fmt.Println(a + b + c)
	stringMain()
}

func stringMain() {
	var x string
	var count int
	fmt.Scan(&x)
	for _, char := range x {
		char := string(char)            // char is rune type at this point, so we need convert to string
		number, _ := strconv.Atoi(char) // atoi return tuple -> (int, error)
		count += number
	}
	fmt.Println(count)
}
