/*
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.
*/

/*
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := strconv.FormatInt(int64(n), 10)
	fmt.Println(string([]rune(s)[0]))
}
*/

// ---------------------------------------------------------------------
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(getFirstNumber(n))
}

func getFirstNumber(n int) int {
	if n < 10 {
		return n
	} else {
		return getFirstNumber(n / 10)
	}
}
