package main

import "fmt"

// СТРУКТУРЫ (2) ВСТРАИВАНИЕ

// Встраивание - когда встраиваем определенное поведение или свойство типа определенное поле структуры
// При наследовании - наследуемый тип получает возможности своего родителя и так же имеет тип родителя т.е. Тип В является подтипом А
// При встраивании - работа по принципу композиции т.е. Тип А содержит тип В. Мы расширяем поведение типа А за счет встраивания
// в него типа В. Но тип А так и остается типом А

type Person struct {
	Name string
	Age  int
}

// Есть структура, человек с именем и возрастом

type WorkExperience struct {
	Name string
	Age  int
}

//(!3) Добавим в woodbuilder эту труктуру

func (p Person) printName() {
	fmt.Println(p.Name)
}

type WoodBuilder struct {
	Person
	Name string // После (!1) добавили свое имя Name
	WorkExperience
}

// Добавим тип Лесной строитель, в него встраиваем тип Person
func main() {
	explanation()
}

func explanation() {
	//(!2) перепишем т.к. добавили поле Name	//builder := WoodBuilder{Person{Name: "Вадим", Age: 30}}
	//(!3) снова перепишем builder := WoodBuilder{Person{Name: "Вадим", Age: 30}, "Боб"}
	builder := WoodBuilder{
		Person{Name: "Вадим", Age: 30},
		"Боб",
		WorkExperience{Name: "Программист", Age: 4}}
	fmt.Printf("Type %T Value: %#v \n", builder, builder)

	// Синтаксический Сахар
	fmt.Println(builder.Person.Age) // Если мы хотим вывести только возраст
	fmt.Println(builder.Age)        // Либо так еще короче. Go переберает WoodBuilder т.к. собственного поля Age нету он
	// берет его из наследуемого типа Person

	//(!3) Теперь Go не понимает какой Age ему вывести т.к. Есть и в Person и в WorkExpirience. Это называется колизия
	// (Colliding). ПО этому в данном случае нужно писать только полный путь!!!

	builder.printName()
	// (!1)	// Вызовем метод printName у нашего билдера. Этот метод выводит Name у структуры Person
	fmt.Printf(builder.Name) // Тут он выведет боба т.к. в WoodBuilder появилось поле Name

	// Shadowing - Свойство, которое находится ближе к корню затеняют более глубокие свойства. По этому вывел Боб а не Вадим
}
