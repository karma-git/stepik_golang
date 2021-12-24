/*
По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

Входные данные
Вводится натуральное число.

Выходные данные
Выведите ответ на задачу.

Sample Input:
50

Sample Output:
1 2 4 8 16 32

*/
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var n int
	var answer string
	fmt.Scan(&n)
	for i := 0; ; i++ {
		x := int(math.Pow(2, float64(i)))
		if x <= n {
			x := strconv.Itoa(x)
			answer += x + " "
		} else {
			break
		}
	}
	fmt.Printf(answer)
}

func mainForum() {
	var mas int
	fmt.Scan(&mas)
	for i := 1; i <= mas; i *= 2 {
		fmt.Print(i, " ")
	}
}
