package main

import "fmt"

func FilterEven(numbers int) {
	switch {
	case numbers%2 == 0:
		fmt.Println(numbers)
	}

}

func main() {

	massiv := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, value := range massiv {
		FilterEven(value)
	}

}
