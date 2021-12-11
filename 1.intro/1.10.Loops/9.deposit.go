/*
Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается. Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.

Входные данные

Программа получает на вход три натуральных числа: x, p, y.

Выходные данные

Программа должна вывести одно целое число.
*/
package main

import "fmt"

func main() {
	var x, p, y, years int
	fmt.Scan(&x, &p, &y)
	for x < y {
		// + debug
		// fmt.Printf("Current year is <%v>\n", years)
		// fmt.Printf("Total money before increase of this year is <%v>\n", x)
		diff := x * p / 100
		// fmt.Printf("Diff in this yeart <%v>\n", diff)
		x += diff
		// fmt.Printf("Money after increase <%v>\n---\n", x)
		years++
	}
	fmt.Println(years)
}

func mainForum() {
	// ref: https://stepik.org/lesson/228263/step/9?discussion=1291993&thread=solutions&unit=200796
	var x, p, y, n int
	fmt.Scan(&x, &p, &y)
	for n = 0; x <= y; n++ {
		x += x * p / 100
	}
	fmt.Println(n)
}
