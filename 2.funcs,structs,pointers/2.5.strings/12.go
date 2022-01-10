/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
Длина пароля должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского алфавита.
На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input:
fdsghdfgjsdDD1
Sample Output:
Ok

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func isDigitOrLetter(x rune) bool {
	if unicode.IsDigit(x) || unicode.Is(unicode.Latin, x) {
		return true
	} else {
		return false
	}
}

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)
	if len(text) >= 5 {
		for _, c := range text {
			if !(isDigitOrLetter(c)) {
				fmt.Println("Wrong password")
				return
			}
		}
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}

func mainForum() {
	a, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]{5,}$", a)
	if matched {
		fmt.Print("Ok")
	} else {
		fmt.Print("Wrong password")
	}
}
