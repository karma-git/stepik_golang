/*
Дается строка. Нужно удалить все символы,
которые встречаются более одного раза и вывести получившуюся строку

Sample Input:
zaabcbd
Sample Output:
zcd
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	for _, v := range s {
		v := string(v)
		if !(strings.Count(s, v) > 1) {
			fmt.Print(v)
		}
	}
}
