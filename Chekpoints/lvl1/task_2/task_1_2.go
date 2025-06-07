package main

import "fmt"

func main() {

	// Решение используя if
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 { // Сначала отдаем приоритет пересекающимся значениям (15, 30, 45 и т.д.)
			fmt.Println("FizzBuzz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}

		fmt.Println(i)
	}

	// Решение используя switch case
	for j := 1; j <= 100; j++ {
		switch {
		case j%3 == 0 && j%5 == 0:
			fmt.Println("FizzBuzz")
		case j%3 == 0:
			fmt.Println("Fizz")
		case j%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(j)
		}
	}
}
