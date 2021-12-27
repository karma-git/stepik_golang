package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Printf("You said <%v>\n", s)
	b := &s
	*b += " World!"
	fmt.Println(b)
	fmt.Println(s)
}
