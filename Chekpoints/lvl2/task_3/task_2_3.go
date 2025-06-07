package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BankAccount struct {
	Balance float64
}

// Метод внесения средств на счет
func (dep *BankAccount) Deposit(raschetDep float64) float64 {
	raschet := dep.Balance + raschetDep
	dep.Balance = raschet
	fmt.Printf("Средства внесены! На вашем балансе: %.2f \n", dep.Balance)
	return dep.Balance
}

// Метод снятия средств со счета
func (with *BankAccount) Withdraw(raschetWith float64) float64 {
	if raschetWith > with.Balance {
		fmt.Printf("На вашем балансе не достаточно средств, ваш баланс:%.2f \n", with.Balance)
		return with.Balance
	}
	raschet := with.Balance - raschetWith
	with.Balance = raschet
	fmt.Printf("Средства сняты! Ваш баранс:%.2f \n", with.Balance)
	return with.Balance
}

// Метод вызова баланса
func (bal BankAccount) ShowBalance() float64 {
	fmt.Printf("Ваш баланс:%.2f \n", bal.Balance)
	return bal.Balance
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	balance := BankAccount{0.00}

	for {
		fmt.Println("Введите что бы выполнить соответствующую операцию:")
		fmt.Println("1 - Внести средства на счет")
		fmt.Println("2 - Снять средства со счета")
		fmt.Println("3 - Посмотреть баланс")
		fmt.Println("4 - Выйти")

		operation, _ := reader.ReadString('\n')  // Читаем строку до символа новой строки (\n, когда нажимаешь Enter)
		operation = strings.TrimSpace(operation) // Убираем лишние пробелы, переносы строк и табуляции в начале и конце строки.
		// Без этого кода почему то не читает

		numOperation, err := strconv.Atoi(operation) // Преобразуем строку в целое число (int) и помещаем в переменную numOperation иначе "1/n"
		if err != nil {                              // Проверяем если введено некорре
			fmt.Println("Ошибка ввода")
			continue
		}

		switch numOperation {
		case 1: // Внесение средств
			fmt.Println("Ведите сумму которую хотите внести:")
			dep, _ := reader.ReadString('\n')
			dep = strings.TrimSpace(dep)
			raschetDep, err := strconv.ParseFloat(dep, 64)
			if err != nil || raschetDep <= 0 { // Проверяем если введено некорре
				fmt.Println("Ошибка ввода")
				continue
			}
			balance.Deposit(raschetDep)

		case 2: // Снятие средств
			fmt.Println("Ведите сумму которую хотите снять:")
			with, _ := reader.ReadString('\n')
			with = strings.TrimSpace(with)
			raschetWith, err := strconv.ParseFloat(with, 64)
			if err != nil || raschetWith <= 0 { // Проверяем если введено некорре
				fmt.Println("Ошибка ввода")
				continue
			}
			balance.Withdraw(raschetWith)

		case 3: // Баланс
			balance.ShowBalance()

		case 4: // Выход
			fmt.Println("Выход")
			return

		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
