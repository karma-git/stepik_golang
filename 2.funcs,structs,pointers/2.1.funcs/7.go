/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные
Вводится четыре числа.

Выходные данные
Необходимо вернуть из функции наименьшее из 4-х данных чисел

Sample Input:
4 5 6 7

Sample Output:
4


*/
package main

import "fmt"

func minimumFromFour() int {
	var m int
	arr := make([]int, 4)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	for i, e := range arr {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}

func main() {
	fmt.Println(minimumFromFour())
}
