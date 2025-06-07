package main

import "fmt"

// Интерфейс - специальный тип в Go, который является набором методов, которые дожен реализовать этот тип

type Runner interface {
	Run() string
}

// Задаем интерфейс Runner и в фигурных скобках перечесляем наши методы

type Swimmer interface {
	Swim() string
}
type Flyer interface {
	Fly() string
}

// Значение интерфейсного типа более сложное и состоит из двух частей:
// 1 - знание о конкретном типе
// 2 - знание о значении конкретного типа

type Human struct {
	Name string
}

// Создаем структуру с именем человека, далее для нее обьявляем метот Run который

func (h Human) Run() string {
	return fmt.Sprintf("Человек %s бегает", h.Name)
}

// Зеленый значек справа обозначает что мы заимплементировали интерфейс, т.е. выполнили все его методы

func main() {

	var runner Runner

	var unnamedRunner *Human
	fmt.Printf("Tipe: %T Value: %#v \n", unnamedRunner, unnamedRunner)

	runner = unnamedRunner
	fmt.Printf("Tipe: %T Value: %#v \n", runner, runner)
	if runner == nil {
		fmt.Println("Runner is nil")
	}
	// Теперь интерфейсное значение не nil, оно содержит знание о конкретном типе Human и о конкретном значении имени Name
	// Но это значение пока что не задано, по этому выводится nil

	nameRunner := &Human{Name: "Вадим"}

	runner = nameRunner
	fmt.Printf("Tipe: %T Value: %#v \n", runner, runner)
	// Теперь интерфест содержит и знание о конкретном типе Human и о его значении Name которое не nill а Вадим
}
