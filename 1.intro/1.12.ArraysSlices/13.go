/*
Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4), а затем NN целых чисел-элементов среза.
На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
Sample Input:
5
41 -231 24 49 6
Sample Output:
49
*/
package main

import "fmt"

func main() {
	var n, el int
	var arr []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&el)
		arr = append(arr, el)
	}
	fmt.Println(arr[3])
}

func mainForum() {
	// ref: https://stepik.org/lesson/228265/step/13?discussion=1132788&thread=solutions&unit=200798
	var n int

	fmt.Scan(&n)
	m := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m[i])
	}
	fmt.Println(m[3])
}
