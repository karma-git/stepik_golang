/*
Заданы три числа - a, b, c (a < b < c)a,b,c(a<b<c) - длины сторон треугольника.
Нужно проверить, является ли треугольник прямоугольным.
Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"
*/
package main

import "fmt"

func triangleIsRight(max, x, y int) {
	if max*max == x*x+y*y {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	switch {
	case a > b && a > c:
		triangleIsRight(a, b, c)
	case b > a && b > c:
		triangleIsRight(b, a, c)
	case c > a && c > b:
		triangleIsRight(c, a, b)
	}
}
