package main

import "fmt"

func main() {
	getSlise()
	copySlice()
	deleteElement()
}

// ПОЛУЧЕНИЕ СЛАЙСА ИЗ МАССИВА. ПОЛУЧЕНИЕ СЛАЙСА ИЗ СЛАЙСА

func getSlise() {
	intArr := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Type: %T Value: %#v \n", intArr, intArr)

	intSlice := intArr[1:3] // Создаем слайт который будет содержать массив intArr, но не весь
	// Будет сожержать массив от элемента с индексом 1 до элемента с индексом 3 (!!!но не включая элемент с индексом 3)
	fmt.Printf("Type: %T Value: %#v \n", intSlice, intSlice)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(intSlice), cap(intSlice))
	// Капасити рассчитывает от элемента который мы указали первым до последнего индекса intArr

	fullSlice := intArr[:] // Создали слайс со всеми элементами массива intArr
	fmt.Printf("Type: %T Value: %#v \n", fullSlice, fullSlice)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(fullSlice), cap(fullSlice))

	// Получим слайс на основе слайса fullSlice
	sliseFromSlice := fullSlice[:3] // Указываем диапазон от начального значения до элемента с индексом 3
	// (так же элемент с индексом 3 не включается)
	fmt.Printf("Type: %T Value: %#v \n", sliseFromSlice, sliseFromSlice)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(sliseFromSlice), cap(sliseFromSlice))

	// Теперь заменим элемент в исходном массиве, на который ссылаются все наши созданные слайсы
	intArr[2] = 500
	// В каждом созданном нами слайсе поменяетс язначение т.к. они копируют указатель от массива intArr
	fmt.Printf("Type: %T Value: %#v \n", intArr, intArr)
	fmt.Printf("Type: %T Value: %#v \n", intSlice, intSlice)
	fmt.Printf("Type: %T Value: %#v \n", fullSlice, fullSlice)
	fmt.Printf("Type: %T Value: %#v \n", sliseFromSlice, sliseFromSlice)
}

// КОПИРОВАНИЕ СЛАЙСА
func copySlice() {
	destination := make([]string, 0, 2)           // Создадим слайс с количеством значений 0 и вместимостью 2
	source := []string{"Vasya", "Petya", "Katya"} // Создадим второй слайс с тремя значениями

	// В го есть функция copy, она копирует, указывается элемент слайс в который мы хотим копировать и затем
	// слайс из которого хотим копировать
	fmt.Println("Copied", copy(destination, source))
	fmt.Printf("Type: %T Value: %#v \n", destination, destination)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(destination), cap(destination))
	// Однако в destination мы указали количество значений 0 - у нас нихера не скопируется.
	// т.к. Копируется минимальное количество элментов которое есть в одном из слайсов

	destination = make([]string, 2, 3) // Теперь задаем количество элементов 2 а капасити 3
	fmt.Println("Copied", copy(destination, source))
	fmt.Printf("Type: %T Value: %#v \n", destination, destination)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(destination), cap(destination))
	// Скопируется именно 2 элемента т.к. в одном из слайсов (destination) минимум 2 элемента

	destination = make([]string, len(source)) // Теперь зададим количество элементов равное количеству элементов в слайсе source
	fmt.Println("Copied", copy(destination, source))
	fmt.Printf("Type: %T Value: %#v \n", destination, destination)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(destination), cap(destination))
	// На этот раз скопировались все элементы

	//!!! В пустой слайс с дефолтным значением (nill) само собой копировать не получется

	// Есть еще фишка если мы хотим скопировать слайс с 0 элементами
	rightCopy := append(make([]string, 0, len(source)), source...)
	// через make создаем слайс с 0 элементами, но с капасити равной длинне слайса source
	// через append - добавляем все значения разложенного на несколько аргументов слайса source (source...)
	fmt.Printf("Type: %T Value: %#v \n", rightCopy, rightCopy)
	fmt.Printf("Lenght: %d Capasity: %d \n", len(rightCopy), cap(rightCopy))
}

// УДАЛЕНИЕ ЭЛЕМЕНТА ИЗ СЛАЙСА
func deleteElement() {
	slice := []int{1, 2, 3, 4, 5} // есть некий слайс
	i := 2                        // задаем элемент который хотим удалить (3)

	withAppend := append(slice[:i], slice[i+1:]...) // Попробуем создать новый слайс
	// Мы берем все элементы до i, но не включая i (т.е. значения 1 и 2)
	// Затем добавляем все элементы которые начинаются с i+1
	// Тем самым мы создали слайс, но исключили элемент i
	fmt.Printf("Type: %T Value: %#v \n", withAppend, withAppend)
	// !!!Однако выполняя такую операцию мы ломаем исходный слайс
	fmt.Println(slice)
	// Что бы этого избежать нужно не создавать новый слайс, а переприсваивать старый
	slice = append(slice[:i], slice[i+1:]...)
	fmt.Printf("Type: %T Value: %#v \n", slice, slice)
	// !!!Теперь все правильно

	// Есть второй способ, но это пиздец какой-то геморойный
	slice = []int{1, 2, 3, 4, 5} // тот же слайс

	slice = slice[:i+copy(slice[i:], slice[i+1:])] // !!!Так же не создаем новый слайс, а меняем текущий
	// Мы берем к элементу i прибавляем количество скопированных элементов
	// Копирование работает так - в качестве начального слайса берем от i До конца (3, 4, 5)
	// И мы в эту часть копируем i+1 до конца (4, 5)
	// т.е. в значение 3, 4, 5 скопируется 4, 5 (вместо 3 - 4, вместо 4 - 5 и в конце останется 5)
	// т.е. вид будет типа таким 1, 2, 4, 5, 5
	// Но при копировании мы взяли минимальное количество элементов два и прибавили к i. По этому скопируются только 4 и 5
	// кароч я ебу формула :D
	fmt.Printf("Type: %T Value: %#v \n", slice, slice)
}
