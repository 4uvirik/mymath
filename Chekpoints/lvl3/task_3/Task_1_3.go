package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите число, а я посчитаю сумму его цифр:")
	digitsString, _ := reader.ReadString('\n')
	digitsString = strings.TrimSpace(digitsString)
	digits, err := strconv.Atoi(digitsString)
	if err != nil {
		fmt.Println("Введить нужно целое число епта! При помощи цифр... ну там 1,6,8 и т.д.")
		return
	}

	total := sumDigits(digits)
	fmt.Println("Сумма цифр числа", digits, "равна", total)

}

func sumDigits(digits int) int {
	var SumDigits = 0
	for digits > 0 { // Вводим новую переменную, что бы область видимости была только внутри цикла (иначе при выводе "а" будет 0)
		SumDigits += digits % 10 // отсекаем остаток (правую цифру) и прибавляем его к следующему отсеченному остатку в цикле
		digits = digits / 10     // каждый цикл уменьшает число убрав правую цифру
	}
	return SumDigits
}
