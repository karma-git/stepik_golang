/*
	Часовая стрелка повернулась с начала суток на d градусов.
	Определите, сколько сейчас целых часов h и целых минут m.
*/

package main

import "fmt"

func main() {
	var a uint16
	fmt.Scan(&a)

	// In our hour, offset of hour arrow == 360 / 12
	currHour := a / 30
	// Then, value of 1 degree == 60 min / 30
	currMin := 2 * (a % 30)
	fmt.Printf("It is %v hours %v minutes.\n", currHour, currMin)
}
