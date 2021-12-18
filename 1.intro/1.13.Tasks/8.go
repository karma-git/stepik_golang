/*
По данным числам, определите количество чисел, которые равны нулю.
Входные данные
Вводится натуральное число N, а затем N чисел.
Выходные данные
Выведите количество чисел, которые равны нулю.
*/
package main

import "fmt"

func main() {
	var n, x, counter int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x == 0 {
			counter++
		}
	}
	fmt.Println(counter)
}
