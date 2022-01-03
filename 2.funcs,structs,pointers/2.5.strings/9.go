package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, substr string
	fmt.Scan(&str, &substr)
	fmt.Print(strings.Index(str, substr))
}
