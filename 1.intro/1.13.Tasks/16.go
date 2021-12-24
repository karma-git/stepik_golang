/*
Из натурального числа удалить заданную цифру.

Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.

Выходные данные
Вывести число без заданных цифр.

Sample Input:
38012732
3

Sample Output:
801272
*/
package main

import "fmt"

func main() {
	var str, target string
	fmt.Scan(&str, &target)
	for _, char := range str {
		char := string(char)
		if char != target {
			fmt.Print(char)
		}
	}
}