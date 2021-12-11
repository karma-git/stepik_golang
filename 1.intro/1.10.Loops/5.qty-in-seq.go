/*
Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.

Формат входных данных

Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0 (само число 0 в последовательность не входит, а служит как признак ее окончания).
*/
package main

import "fmt"

func main() {
	var max, n, count int
loop:
	for {
		fmt.Scan(&n)
		switch {
		case n == 0:
			break loop
		case n > max:
			max, count = n, 1
		case n == max:
			count++
		}
	}
	fmt.Println(count)
}

func mainForum() {
	var a, s, p int

	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		if a > s {
			s = a
			p = 0
		}
		if a == s {
			p++
		}
	}
	fmt.Println(p)
}
