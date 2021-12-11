/*
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.
*/
package main

import "fmt"

func isMultipleOfEight(n int) bool {
	if n%8 == 0 {
		return true
	} else {
		return false
	}
}

func isNumberHasTwoDigits(n int) bool {
	if n/100 == 0 && !(n/10 == 0) {
		return true
	} else {
		return false
	}
}

func main() {
	var seq, sum int
	fmt.Scan(&seq)
	for i := 1; i < seq+1; i++ {
		var n int
		fmt.Scan(&n)
		if isMultipleOfEight(n) && isNumberHasTwoDigits(n) {
			sum += n
		}
	}
	fmt.Println(sum)
}

/*
ref: https://stepik.org/lesson/228263/step/4?discussion=1623003&thread=solutions&unit=200796
package main

import "fmt"

func main() {
	var n, sum int
	for fmt.Scan(&n); n > 0; n-- {
		var x int
		if fmt.Scan(&x); x > 9 && x < 100 && x%8 == 0 {
			sum += x
		}
	}
	fmt.Println(sum)
}
*/
