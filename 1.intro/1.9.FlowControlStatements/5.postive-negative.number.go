/*
На ввод подается целое число.
Если число положительное - вывести сообщение "Число положительное", если число отрицательное - "Число отрицательное". Если подается ноль - вывести сообщение "Ноль".
Выводить сообщение без кавычек.
*/

package main

import "fmt"

func main() {
	var d int32
	fmt.Scan(&d)

	switch d {
	case 0 < d:
		fmt.Println("Число отрицательное")
	case 0 > d:
		fmt.Println("Число положительное")
	case 0 == d:
		fmt.Println("Ноль")
	}

}
