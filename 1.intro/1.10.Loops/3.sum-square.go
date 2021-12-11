/*
Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B).
Вывести сумму всех чисел от A до B  включительно.
*/
package main

import "fmt"

func main() {
	var a, b, sum int
	fmt.Scan(&a, &b)

	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Println(sum)
}

/*
Арифметическая прогрессия
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Println((b - a + 1) * (a + b) / 2)

}
*/
