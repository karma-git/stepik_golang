/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму.

Пример вызова вашей функции:

a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1

Sample Input:
*/
package main

import "fmt"

func sumInt(a ...int) (c, s int) {
	for _, el := range a {
		c++
		s += el
	}
	return
}

func main() {
	a, b := sumInt(1, 2, 3, 4, 5)
	fmt.Println(a, b)
}
