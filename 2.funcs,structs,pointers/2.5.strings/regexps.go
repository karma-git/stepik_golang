package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegexp := regexp.MustCompile(`^[0-9]{5,6}$`)

	match5 := sampleRegexp.MatchString("84941")
	match6 := sampleRegexp.MatchString("184941")

	match4 := sampleRegexp.MatchString("1234")
	match7 := sampleRegexp.MatchString("1234567")

	fmt.Printf("For 5: %t\n", match5)
	fmt.Printf("For 6: %t\n", match6)

	fmt.Printf("For 4: %t\n", match4)
	fmt.Printf("For 7: %t\n", match7)
}
