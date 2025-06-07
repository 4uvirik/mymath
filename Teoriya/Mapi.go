package main

import (
	"fmt"
)

// Мапы (карты, словари) - Тип данных созданный на основе хэш таблицы
// Хэш таблица - тип данных в которой значения хранятся в виде пар - Ключ: Значение
// Ключ дллжен быть comparable (сравниваемый с другими типами)
// Incomparable types (Несравниваемые типы):
// Slices, maps, functions, struct с incomparable типами полями

type User struct {
	Id   int64
	Name string
}

func main() {

	var defaultMap map[int64]string // в квадратных скобках тип ключа, затем тип значение

	fmt.Printf("Type: %T Value: %#v \n", defaultMap, defaultMap)
	fmt.Printf("Len: %d \n\n", len(defaultMap))
	// Дефолтное значение - nil

	// Создадим мапу через функцию make
	mapByMake := make(map[string]string)
	fmt.Printf("Type: %T Value: %#v \n", mapByMake, mapByMake)

	// Создадим такую же мапу, но с указаным количеством элементов которое мы хотели бы разместить в мапе
	mapByMakeWithCap := make(map[string]string)
	fmt.Printf("Type: %T Value: %#v \n", mapByMakeWithCap, mapByMakeWithCap)

	// Создадим мапу с двумя значениями
	mapByLiberal := map[string]int{"Vasya": 18, "Dima": 20}
	// "Vasya": 18 - это и есть Ключ: Значение
	fmt.Printf("Type: %T Value: %#v \n", mapByLiberal, mapByLiberal)
	fmt.Printf("Len: %d \n\n", len(mapByLiberal))

	// Четвертый способ создание мапы через функцию new
	mapWithNew := *new(map[string]string)
	// Мы получаем указатель на мапы, что бы сразу разименовать пишем перед new - *
	fmt.Printf("Type: %T Value: %#v \n", mapWithNew, mapWithNew)

	// ВСТАВКА И ОБНОВЛЕНИЕ ЗНАЧЕНИЙ МАПЫ

	mapByMake["First"] = "Vasya"
	// Мы в мапу mapByMake в включ "First" положили новое значение "Vasya"
	// Вставка работает в том случае если не было ключа в мапе
	fmt.Printf("Type: %T Value: %#v \n", mapByMake, mapByMake)
	fmt.Printf("Len: %d \n\n", len(mapByMake))

	// Обновим значение ключа "First"
	mapByMake["First"] = "Petya"
	// Та же запись что и выше, но когда ключ есть Go просто обновит значение
	fmt.Printf("Type: %T Value: %#v \n", mapByMake, mapByMake)
	fmt.Printf("Len: %d \n\n", len(mapByMake))

	// Получим и выведем значение мапы
	fmt.Println(mapByLiberal["Vasya"])
	// Ключ "Vasya" у мапы mapByLiberal есть и его значение равно 18

	// Попробуем вывести значение не существующего ключа
	fmt.Println(mapByLiberal["Second"])
	// т.к. Ключа такого нету мы получим дефолтное значение типа (int) 0

	// Но что если в мапе был ключ "Second" и в нем лежало значение 0, нужно различать дефолтное значение от значения 0
	// В Go есть расширение получения значения из мапы
	value, ok := mapByLiberal["Second"]
	// Задаем две переменные. Первая - значение которое мы получим. Второе - признак существует значение или нет
	fmt.Printf("Value: %d InExist: %t\n\n", value, ok)
	// Мы получаем значение value - 0 но, оно может быть и дефолтным. Сущесрвует ли ключь в мапе - false

	// УДАЛЕНИЕ ЗНАЧЕНИЯ ИЗ МАПЫ

	// Удаляем при помощью функции delete
	delete(mapByMake, "First")
	fmt.Printf("Type: %T Value: %#v \n", mapByMake, mapByMake)

	// ИТЕРАЦИИ ПО МАПАМ

	// Создадим какую то мапу с четырьмя значениями
	mapForIteraion := map[string]int{"First": 1, "Second": 2, "Third": 3, "Fourth": 4}

	// Используя цикл for ranfe посмотрим ключи и значения
	for key, val := range mapForIteraion { // Создаем две переменные от мапы
		fmt.Printf("Key: %s Value: %d\n", key, val) // Выводим в одной ключь, во второй значение
	}
	// При переборе мы получаем разные значения, самими разработчиками Go не гарантировано повторяемость последовательности
	// т.е. Последовательности при переборе мапы нет!

	// КАК ИСПОЛЬЗОВАТЬ МАПЫ

	// Например как набор уникальных сущностей
	// Предположим мы собираем пользователей с нескольких сервисов и на выходе получаем следующий слайс с пользователями
	users := []User{
		{
			Id:   1,
			Name: "Vasya",
		},
		{
			Id:   45,
			Name: "Petya",
		},
		{
			Id:   57,
			Name: "John",
		},
		{
			Id:   45,
			Name: "Petya",
		},
	}
	// Мы видим что в данном слайсе дважды встречается Petya, А нам надо создать список УНИКАЛЬНЫХ пользователей
	uniqueUsers := make(map[int64]struct{}, len(users))
	// Создаем мапу с ключем int64 (сюда складываем Id пользователя) и в качестве значения будем использовать
	// пустую структуру (она не занимает нисколько памяти). И в качстве длины используем длину нашего слайса users
	for _, user := range users { // Мы итерируемся по слайсу users
		if _, ok := uniqueUsers[user.Id]; !ok { // Мы проверяем, значение user по ключию Id,
			// если в нашей мапе ключ ранвый Id нашего user не найден, то мы туда его помещаем
			uniqueUsers[user.Id] = struct{}{} // В нашу мапу добавляем значение, в качестве ключа используем Id,
			// А в качестве значения пустую структуру
		}
	}
	fmt.Printf("Type: %T Value: %#v \n", uniqueUsers, uniqueUsers)

	// !!! (1)
	fmt.Println(findInSlice(57, users))

	// !!! (2)
	// Что бы не делать слишком много итераций есть лайфхак
	// Создатим мапу для нашего слайса
	usersMap := make(map[int64]User, len(users))
	for _, user := range users {
		if _, ok := usersMap[user.Id]; !ok {
			usersMap[user.Id] = user
		}
	}

	fmt.Println(findInMap(57, usersMap))
}

// Второй пример использования мапы - быстрый поиск значения
// Используем тот же самый слайс с пользователями
// Попробуем найсти пользователя в слайсе !!!(1)
func findInSlice(id int64, users []User) *User { // Принимаем id который нам нужно найсти и принимаем слайс users
	for _, user := range users { // Что бы найти пользователя итерируемся по всем нашим юзерам
		if user.Id == id { // и когда находим юзра у которого Id будет такой же какой мы запросили (id)
			return &user // То мы возвращаем указатель на нашего юзера
		}
	}
	return nil // Если мы не нашли нужное нам значение (Id == id) то мы возвращаем nill
	// В качестве возвращаемого значения используем указатель на юзера (*User)
}

// В данном случае, если пользователей 1000, нам нужно искать для каждого id, и совершать несколько итераций
// это очень времязатратно

func findInMap(id int64, usersMap map[int64]User) *User {
	if user, ok := usersMap[id]; ok {
		return &user
	}
	return nil
}
