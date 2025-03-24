package hangman

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	letter := getLetter()
	fmt.Println(letter)
}

func getLetter() string {
	var letter string
	for {
		fmt.Println("Введите одну букву на русском языке:")
		fmt.Scan(&letter)
		if utf8.RuneCountInString(letter) == 1 && checkString(letter) {
			return strings.ToLower(letter)
		}
	}
}

func checkString(str string) bool {
	re, _ := regexp.Compile("[А-Яа-яЁё]")
	return re.MatchString(str)
}
