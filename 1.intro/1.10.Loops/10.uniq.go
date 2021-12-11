/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.

Входные данные

Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.

Выходные данные

Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения в первом числе.

Внимание! Если вам сложно решить эту задачу, пропустите и изучайте материал дальше, затем вернетесь к этой задаче позже.

Sample Input:

564 8954
Sample Output:

5 4
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b, container string
	fmt.Scan(&a, &b)
	for _, char := range a {
		char := string(char)
		if strings.Contains(b, char) && !(strings.Contains(container, char)) {
			container += fmt.Sprintf("%v ", char)
		}
	}
	fmt.Println(container)
}

func mainForum() {
	// ref: https://stepik.org/lesson/228263/step/10?discussion=1614811&thread=solutions&unit=200796
	var a, b, an, bn int
	fmt.Scan(&a, &b)

	for i := 1000; i >= 1 && a > 0; i = i / 10 {
		if a/i == 0 {
			continue
		}
		an = a / i % 10
		for i := 1; i < 10000 && b > 0; i = i * 10 {
			bn = b / i % 10
			if bn == an {
				fmt.Print(an, " ")
				break
			}
		}
	}
}

// https://stepik.org/lesson/228263/step/10?discussion=1614811&thread=solutions&unit=200796

func matchDigits(x, y int) {
	if x > 10 {
		matchDigits(x/10, y)
		matchDigits(x%10, y)
	} else if y > 10 {
		matchDigits(x, y/10)
		matchDigits(x, y%10)
	} else if x == y {
		fmt.Print(x, " ")
	}
}

func matchForumRecursion() {
	var a, b int

	fmt.Scan(&a, &b)
	matchDigits(a, b)
}
