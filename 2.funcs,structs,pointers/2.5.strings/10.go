/*
На вход дается строка, из нее нужно сделать другую строку,
оставив только нечетные символы (считая с нуля)

Sample Input:
ihgewlqlkot
Sample Output:
hello
*/
package main

import (
	"fmt"
)

func main() {
	var s, result string
	fmt.Scan(&s)
	for i := 0; i < (len(s)); i++ {
		if i%2 == 1 {
			result += string(s[i])
		}
	}
	fmt.Println(result)
}
