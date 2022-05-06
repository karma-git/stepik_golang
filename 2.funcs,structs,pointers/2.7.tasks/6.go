/*
ref: https://stepik.org/lesson/229321/step/7?unit=201907
*/
package main

import (
	"fmt"
	"math"
)

var k, p, v float64 = 1296, 6, 6

func main() {
	fmt.Println(T())
}

func M() (m float64) {
	m = p * v
	return
}

func W() (w float64) {
	w = math.Sqrt(k / M())
	return
}

func T() (t float64) {
	t = 6 / W()
	return
}
