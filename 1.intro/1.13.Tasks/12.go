/*
По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".

Входные данные
Дано число n (0<n<100).

Выходные данные
Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.

Sample Input:
10

Sample Output:
10 korov

*/
package main

import "fmt"

func inRange(x, a, b int) bool {
	for ; a < b; a++ {
		if x == a {
			return true
		}
	}
	return false
}

func main() {
	var n int
	fmt.Scan(&n)

	// fmt.Println(n % 10)

	switch {
	case n%10 == 1 && n != 11:
		fmt.Printf("%v korova", n)
	case inRange(n%10, 5, 10) || n%10 == 0 || inRange(n, 10, 21):
		fmt.Printf("%v korov", n)
	case inRange(n%10, 2, 5) && !(inRange(n, 10, 21)):
		fmt.Printf("%v korovy", n)
	}
}
