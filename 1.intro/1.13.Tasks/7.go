/*
Даны два числа. Найти их среднее арифметическое.

Формат входных данных
На вход дается два целых положительных числа a и b.

Формат выходных данных
Программа должна вывести среднее арифметическое чисел a и b (ответ может быть целым числом или дробным)
*/
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(float64(a+b) / 2)
}