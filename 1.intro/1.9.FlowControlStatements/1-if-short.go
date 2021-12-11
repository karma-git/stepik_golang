package main

import "fmt"

func main() {
	a := 6
	b := 100
	if v := a * b; v < 100 {
		fmt.Printf("Произведение a*b, меньше 100 <%v>\n", v)
	} else {
		fmt.Println("Не попали в условие")
	}
}
