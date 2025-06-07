package main

import "fmt"

// Массив - Тип данных который содержет фиксированное количество элементов определенного типа

type Person2 struct {
	Name string
	Age  int
}

func main() {
	//arrays()  // убрать 2 слеша если нужно посмотреть код массивов
	slices()
}

func arrays() {
	var intArr [3]int                                    // Создаем переменную массив
	fmt.Printf("Type: %T Value: %#v \n", intArr, intArr) // выводим тип массив состоящий из 3 интов дефолтных значениий (0)

	intArr[0] = 5 // [0] - Индекс массива. Индексация всегда начинается с нуля, НЕ С ЕДЕНИЦЫ!!!
	intArr[1] = 6
	fmt.Printf("Type: %T Value: %#v \n", intArr, intArr) // теперь в массиве значения поменялись

	// Рассмотрим сокращенный синтаксис, присваивая сразу значение
	people := [2]Person2{ // Создаем массив состоящий из двух элементов типа структура Person
		{ // Внутри структуры Person в фигурных скобках заполдняем структуру
			Age:  30,
			Name: "Katy",
		},
		{
			Age:  23,
			Name: "John",
		},
	}
	fmt.Printf("Type: %T Value: %#v \n", people, people)

	// Рассмотрим создание массива где не указываем количество элементов, Go вычеслит его автоматически
	stringsArr := [...]string{"Первый", "Второй", "Третий", "Чтвертый"}
	fmt.Printf("Type: %T Value: %#v \n", stringsArr, stringsArr)

	// Рассмотрим вызов встроенных в Go функций для массивов "len" и "cap"
	// len - Длина массива (количество элементов, которые сейчас содержатся в массиве)
	// cap - Вместимость массива (объем, который может поместится в массив)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(stringsArr), cap(stringsArr))
	// По сути всегда len = capб просто нужно знать что есть такие функции

	// Рассмотрим Итерации по массиву
	for i := 0; i < len(stringsArr); i++ { // Обьявляем цикл от 0 (т.к. индексация в массиве начинается с 0)
		// до конечного значения длинны нашего массива
		fmt.Printf("index: %d Value: %s \n", i, stringsArr[i])
	}

	// Для массивов есть лучший вариант for range! Его писать проще. Рассмиотрим его
	for inx, value := range stringsArr { // Обьявляем в цикле две переменные например inx и value
		// Далее присваеваем их к range и указываем нужный массив
		// В первую переменную (в нашем случае inx) всегда приходит индекс!
		// Во вторую переменную (в нашем случае value) всегда приходит значение индекса!
		fmt.Printf("index: %d Value: %s \n", inx, value)
	}

	// Если нам нужно только например значение без индекса, пишем вместо нужной переменной _
	for _, value := range intArr {
		fmt.Printf("Value: %d \n", value)
	}

	// Рассмотрим передачу массива в функцию и изменения его значения
	newIntArr := changeArray(intArr)
	fmt.Printf("Type: %T Value: %#v \n", intArr, intArr)
	fmt.Printf("Type: %T Value: %#v \n", newIntArr, newIntArr)

}

func changeArray(arr [3]int) [3]int { // функция принимает массив с 3 значениями
	arr[2] = 3 // и меняем его последний элемент на значение 3
	return arr
}

// Slices - Слайс (срез) - тип данных который может содержать любое количество элементов типа Т
// Часто слайсы называют динамическими массивами

func slices() {
	var defaultSlice []int // СОздаем слайс с дефолтным значением
	fmt.Printf("Type: %T Value: %#v \n", defaultSlice, defaultSlice)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(defaultSlice), cap(defaultSlice))
	// Дефолтное значение у пустого слайса это nill!

	// Создадим переменную и прировняем ей слайс с двумя строками
	stringSliceLiteral := []string{"First", "Second"}
	fmt.Printf("Type: %T Value: %#v \n", stringSliceLiteral, stringSliceLiteral)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(stringSliceLiteral), cap(stringSliceLiteral))
	// Выведет тип слайс с двумя переменными First и Second
	// максимальная вместимость слайса в данном случае будет так же 2 пока не будет передислоцирована память под этот слайс!!!

	// Рассмотрим пример с заданной длиной слайса и максимальным количеством значений
	sliceByMake := make([]int, 0, 5) // Мы создали слайт где 0 - длина слайса, а 5 - Капасити слайса
	// make - специальная функция в Go для слайсов
	fmt.Printf("Type: %T Value: %#v \n", sliceByMake, sliceByMake)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(sliceByMake), cap(sliceByMake))

	// Слайс не хранит данные на прямую! Внутри слайса лежит указатель на массив
	sliceByMake = append(sliceByMake, 1, 2, 3, 4, 5) // Функция append добавляет элемент в наш слайс
	// Тут append принимает изначальный sliceByMake, затем возвращает этот слайс где есть исходные значения и добавленные
	fmt.Printf("Type: %T Value: %#v \n", sliceByMake, sliceByMake)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(sliceByMake), cap(sliceByMake))
	// Теперь  унашего слайса и длинна 5 и капасити 5

	sliceByMake = append(sliceByMake, 6) // Попробуем добавить шестой элемент нашему слайсу
	fmt.Printf("Type: %T Value: %#v \n", sliceByMake, sliceByMake)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(sliceByMake), cap(sliceByMake))
	// Теперь Go видит что мы добавляем шестой элемент в слайс где максимум 5 элементов. И Go выделит в памяти новый массив
	// - у которого емкость будет в 2 раза больше чем у текущего слайса. И наш новый слайс будет ссылвать уже на новый массив в памяти

	// Рассмотрим фукнции for range для слайсов
	for ind, value := range sliceByMake { // Все так же как и для массива
		fmt.Printf("index: %d Value: %s \n", ind, value)
	}

}
