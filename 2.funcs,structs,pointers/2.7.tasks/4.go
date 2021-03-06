/*
Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.

Входные данные

Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только десятичные цифры.

Выходные данные

Выведите максимальную цифру, которая встречается во введенной строке.

Sample Input:

1112221112
Sample Output:

2
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	var max int
	fmt.Scan(&s)
	for _, r := range s {
		char := string(r)
		number, _ := strconv.Atoi(char)
		if number > max {
			max = number
		}
	}
	fmt.Println(max)
}
