/*
Количество минимумов
Найдите количество минимальных элементов в последовательности.

Входные данные

Вводится натуральное число N, а затем N целых чисел последовательности.

Выходные данные

Выведите количество минимальных элементов последовательности.
*/
package main

import "fmt"

func arrMin(a []int) int {
	var m int
	for i, e := range a {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}

func countInArr(a []int, m int) int {
	var counter int
	for _, el := range a {
		if el == m {
			counter++
		}
	}
	return counter
}

func main() {
	// создается массив из чисел
	// находится минимальный элемент массива
	// проверяется количество этих элементов в массиве
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	m := arrMin(arr)
	count := countInArr(arr, m)
	fmt.Println(count)
}
