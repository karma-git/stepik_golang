/*
Поменяйте местами значения переменных на которые ссылаются указатели. После этого переменные нужно вывести.
ВАЖНО: Считайте что пакет main уже объявлен, а также функция main() уже есть.

Sample Input:
2 4

Sample Output:
4 2
*/
package main

import (
	"fmt"
)

func test(x1, x2 *int) {
	tmp := *x1
	*x1, *x2 = *x2, tmp
	fmt.Print(*x1, " ", *x2)
}

func main() {
	var a, b int
	var x, y *int
	fmt.Scan(&a, &b)
	x, y = &a, &b
	test(x, y)
}
