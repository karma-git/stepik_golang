/*
Требуется определить, является ли данный год високосным, напомним:
Год является високосным если он соответствует хотя бы одному из нижеперечисленных условий:
- кратен 400;
- кратен 4, но не кратен 100.
*/

/*
MY SOLUTION
package main

import "fmt"

func main() {
	var y int
	fmt.Scan(&y)

	switch {
	case y%400 == 0:
		fmt.Println("YES")
	case y%4 == 0 && y%100 != 0:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
*/

// FORUM
package main

import "fmt"

func main() {
	var y uint
	fmt.Scan(&y)
	if y%400 == 0 || (y%4 == 0 && y%100 != 0) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
