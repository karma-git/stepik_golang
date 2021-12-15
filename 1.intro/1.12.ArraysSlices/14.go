/*
На ввод подаются пять целых чисел, которые записываются в массив. Однако эта часть программы уже написана.
Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.
*/
package main

import "fmt"

func main() {
	arr := make([]int, 5)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println(max)
}
