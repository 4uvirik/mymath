package main

import (
	"1_Settings/Chekpoints/lvl4/task_2/mymath"
	"fmt"
)

func main() {
	num := []float64{1.11, 2.22, 3.33, 4.44, 5.55}
	result := mymath.Average(num)
	fmt.Printf("Среднее значение для %v равно: %.2f", num, result)
}
