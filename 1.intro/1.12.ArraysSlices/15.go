/*
Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0. Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).
Входные данные
Сначала задано число NN — количество элементов в массиве (1 \leq N \leq 1001≤N≤100). Далее через пробел записаны NN чисел — элементы массива. Массив состоит из целых чисел.
Выходные данные
Необходимо вывести все элементы массива с чётными номерами.
*/
package main

import "fmt"

func main() {
	var n int
	var result string
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if i%2 == 0 {
			result += fmt.Sprintf("%v ", x)
			// or better use:
			// fmt.Print(x, " ")
		}
	}
	fmt.Println(result)
}
