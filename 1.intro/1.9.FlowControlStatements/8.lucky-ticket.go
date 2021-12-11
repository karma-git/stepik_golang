/*
Определите является ли билет счастливым.
Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	firstHalfSum := threeDigSum(n / 1000)
	lastHalfSum := threeDigSum(n % 1000)
	if firstHalfSum == lastHalfSum {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func threeDigSum(n int) int {
	first := n / 100
	second := n / 10 % 10
	third := n % 10
	return first + second + third
}

/*
Держите алгоритм, по которому можно найти каждую цифру n-значного числа num:

Последняя цифра: (num % 101) // 100;
Предпоследняя цифра: (num % 102) // 101;
Предпредпоследняя цифра: (num % 103) // 102;
.....
Вторая цифра: (num % 10n-1) // 10n-2;
Первая цифра: (num % 10n) // 10n-1
*/
