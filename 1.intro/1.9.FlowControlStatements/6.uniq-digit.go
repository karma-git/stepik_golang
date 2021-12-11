/*
https://stepik.org/lesson/232593/step/6?unit=205068
По данному трехзначному числу определите, все ли его цифры различны.
*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := n / 100
	b := n / 10 % 10
	c := n % 10

	if a != b && b != c && c != a {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
