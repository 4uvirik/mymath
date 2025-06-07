package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите пароль")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	err := ValidatePassword(password)
	if err != nil {
		fmt.Println("Ошибки:", err)
	} else {
		fmt.Println("Пароль введен корректно")
	}
}

func ValidatePassword(pass string) error {

	var capital, number int
	for _, simbol := range pass {
		switch {
		case unicode.IsUpper(simbol):
			capital++
		case unicode.IsDigit(simbol):
			number++
		}
	}

	lenght := utf8.RuneCountInString(pass)

	if lenght < 8 {
		fmt.Println("В пароле необходимо минимум 8 символов")
	}
	if capital < 1 {
		fmt.Println("Необходима минимум одна заглавная буква в пароле")
	}
	if number < 1 {
		fmt.Println("Необходима минимум одна цифра в пароле")
	}
	return nil
}
