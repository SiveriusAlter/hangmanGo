package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	letter := getLetter()
	fmt.Println(letter)
	getWord()
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

func getDictionary() []byte {
	file, err := os.Open("../dictionary.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	dict := make([]byte, 64)

	for {
		_, err := file.Read(dict)
		if err == io.EOF {
			break
		}
	}
	return dict
}

func getWord() {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r1.Intn(100))
}
