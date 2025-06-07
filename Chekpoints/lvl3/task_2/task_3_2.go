package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	text, err := ReadFile("Файл.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("Содержимое файла:\n%s", text)
}

func ReadFile(path string) (string, error) {

	path = "Chekpoints/Файл.txt"
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("Файл не существует")
		}
		return "", fmt.Errorf("Не возможно получить информацию о файле:", err)
	}

	size := info.Size()
	if size > 1024*1024 {
		return "", fmt.Errorf("Файл слишком большой (>1 МБ)")
	}

	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("Ошибка чтения файла:", err)
	}
	defer file.Close()

	var stroka strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stroka.WriteString(scanner.Text())
		stroka.WriteString("\n")
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("Ошибка чтения файла:", err)
	}

	return stroka.String(), nil
}
