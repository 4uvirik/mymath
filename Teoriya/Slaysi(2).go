package main

import (
	"fmt"
	"unsafe"
)

func main() {
	variadicFunctions()
	convertToArrayPointer()
	passToFunction()
	sliceWithNew()
}

func variadicFunctions() {
	showAllElements(1, 2, 3) // Передадим в нашу функцию сначала 3 числа, затем 7
	showAllElements(1, 2, 3, 4, 5, 6, 7)

	// Разберем кейс если есть слайс и у него есть сколько то элементов.
	firstSlice := []int{5, 6, 7, 8}
	// И мы хотим вывести все его элементы в нашей фукнции showAllElements.
	// Но данная функция принимает все агрументы перечисленные через запятую, а не слайс целиком
	// В таком случае в Go можно разложить массив или слайс на отдельные элементы и сразу отправить эти параметры в функцию
	showAllElements(firstSlice...) // Равнозначна тому что мы вызвали функцию и перечислили все элементы (5, 6 ,7 ,8)

	// Рассмотрим пример если к нашему слайсу нужно добавить сразу много элементов через append
	// Это можно сделать создав второй слайс с этими элементами
	secondSlice := []int{9, 3, 1, 5, 2}
	// Далее просто добавляем этот слайс через append следующим кодом
	newSlice := append(firstSlice, secondSlice...)
	fmt.Printf("Type: %T, Value: %#v \n", newSlice, newSlice)
}

func showAllElements(values ...int) { // ... - значит что данная функция можем принять сколько угодно параметров типа Int
	// И все эти аргументы фукнции будут сложены в слайс values
	for _, val := range values {
		fmt.Println("Value:", val)
	}
	fmt.Println()
}

// Рассмотрим структуру слайса, как он выглядит и из чего состоит
type _slice struct {
	elements unsafe.Pointer // - Указатель на массив с определенным типом данных
	len      int            // - Количество элементов
	cap      int            // - Текущая вместимость
}

// Рассмотрим ВОЗМОЖНОСТЬ КОНВЕРТАЦИИ СЛАЙСА В УКАЗАТЕЛЬ НА МАССИВ
func convertToArrayPointer() {
	initialSlice := []int{1, 2} // Есть  слайс с двумя значениями
	fmt.Printf("Type: %T Value: %#v \n", initialSlice, initialSlice)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(initialSlice), cap(initialSlice))
	// Конвертируем в указатель на массив
	intArray := (*[2]int)(initialSlice) // Берем слайс initialSlice и в скобках указываем тип к которому приводим
	// * - указатель
	// [2]int - массив с двумя элементами
	fmt.Printf("Type: %T Value: %#v \n", intArray, intArray)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(intArray), cap(intArray))
	// !!! Это можно сделать только если длина массива совпадает с длиной слайса

}

// ПЕРЕДАЧА СЛАЙСА В ФУНКЦИЮ. ИЗМЕНЕНИЕ ЗНАЧЕНИЯ СЛАЙСА В ФУНКЦИИ ИЛИ ИЗМЕНЕНИЕ КОЛИЧЕСТВО ЭЛЕМЕНТОВ СЛАЙСА В ФУНКЦИИ
func passToFunction() {
	initialSlice := []int{1, 2} // есть стандартный слайс
	fmt.Printf("Type: %T Value: %#v \n", initialSlice, initialSlice)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(initialSlice), cap(initialSlice))
	// !!! Т.к. Слайс внутри содержит ссылку на массив то передавая слайс в функцию его значение копируется.
	// Но вместо того что бы копировать все элементы массива, копируется именно указатель.
	// Передавая слайс в функцию мы передаем указатель на массив, на который ссылается данный слайс

	changeValue(initialSlice)
	fmt.Printf("Type: %T Value: %#v \n", initialSlice, initialSlice)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(initialSlice), cap(initialSlice))
	// Меняется второй элемент на 15

	// Добавим третий элемент нашему слайсу
	newSlice := append(initialSlice, 3)
	fmt.Printf("Type: %T Value: %#v \n", newSlice, newSlice)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(newSlice), cap(newSlice))
	// Создастся новый массив внутри слайса с длиной 3 и вместимостью в два раза больше предыдущей

	newSlice2 := appendValue(newSlice)
	fmt.Printf("Type: %T Value: %#v \n", newSlice, newSlice)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(newSlice), cap(newSlice))
	// Новый массив будет только после того как мы вернем значение в функции appendValue

	fmt.Printf("Type: %T Value: %#v \n", newSlice2, newSlice2)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(newSlice2), cap(newSlice2))
}
func changeValue(slice []int) { // создаем функцию которая меняем значение элемента
	slice[1] = 15 // меняем значение второго элемента на 15
}

func appendValue(slice []int) []int { // Создаем функцию которая добавляет 2 элемента
	slice = append(slice, 4, 5) // Тут появляется ссылка на нвоый массив т.к. превысилось значение элементов
	fmt.Printf("Type: %T Value: %#v \n", slice, slice)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(slice), cap(slice))
	// Если не вернуть значение то массив в слайсе newSlice2 не заменится
	return slice
}

// СОЗДАНИЕ СЛАЙСА ЧЕРЕЗ ФУНКЦИЮ NEW

func sliceWithNew() {
	slicePointer := new([]int) // создаем переменную, вызывает функцию new и показываем что это будет слайс int'ов
	// Функция new возвращает указатель на какой то тип. но значения будут дефолтные (Nill) т.е. и длина и капасити будут 0
	fmt.Printf("Type: %T Value: %#v \n", slicePointer, *slicePointer)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(*slicePointer), cap(*slicePointer))
	// Не забываем что необходимол разыменовывать указатель
	// В данном случае мы никак не можем дополнительно повлиять на наш исходный слайт. Остается долько добавлять элементы

	newSlice2 := append(*slicePointer, 1)
	fmt.Printf("Type: %T Value: %#v \n", newSlice2, newSlice2)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(newSlice2), cap(newSlice2))
	// Теперь элементы есть и с ним можно работать
}
