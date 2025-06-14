package main

import (
	"fmt"
	"runtime"
	"time"
)

// Горутина - Легковестный поток
// Размер стека горутины 2Kb
// Горутинами управляет планировщик Go. В отличии от потоков, которыми управляет операционная систем
// Программист не создает треды сам, работа с ними происходит "под капотом", нам доступны только горутины
// В каждой программе на Go есть как минимум одна горутина - функция main

func main() {

	// Создаем горутину при помощи "go" перед функцией
	go showNumbers(100)

	// !!!(2)
	defer fmt.Println(1)
	defer fmt.Println(2)

	/// !!!(3)
	fmt.Println(sum(2, 3))

	runtime.GOMAXPROCS(1)         // Уменьшаем максимальное количество одновременно выполняющихся горутин до 1
	fmt.Println(runtime.NumCPU()) // Смотрим сколько логических ядер на данном компьютере
	// Мы можем повлиять на количество горутин которые у нас выполняются командой runtime.GOMAXPROCS(1)
	// Однако максимальное количество будет зависеть от количества логических ядер на компе

	// При помощи runtime.Gosched() мы можем переключаться на другую горутину
	// runtime.Gosched()

	// Планировщик Go работает таким образом что сам решает когда какие горутины запускать
	// Например если мы напишем основной горутине спать одну секунду, он логически в эту секунду будет выполнять
	// другую горутину, что бы не терялось время
	time.Sleep(time.Second)

	fmt.Println("Exit")

	makePanic()

}

func showNumbers(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}

// В данном случае не успевает печатать числа т.к. цель планировщика Go завершить программу как можно скорее
// !!! Для того чтобы правильно работать с горутинами их нужно синхронизировать
// Одновременно может выполняться столько горутин сколько логических ядер у вашего процессора
// командой fmt.Println(runtime.NumCPU())

// DEFERRED ФУНКЦИИ
// это функции вызываемые командой deferred, и они не выполняеются сразу, а складываются в специальный стек
// и они начинают выполнятся когда выполняется конструкция return, причем выполняются в обратном порядке
// !!!(2)

func sum(x, y int) (sum int) { // создадим функцию sum которая складывает 2 числа
	defer func() { // в нее добавим deferred функцию, которая умножает сумму на 2
		sum *= 2
	}()
	sum = x + y
	return // в результате после выполнения функции !!!(3) Получим не 5 а 10, т.к. отложенная deferred
	// Выполнила действие после основной функции после return
}

// PANIC И RECOVER
// Панка - особая ситуация в Go когда наша программа говорит о том что произошло что то неожиданное
// и скорее всего дальше работать не сможет

// Панику можно создать при помощи функции panic, внутрь можем положить любое значение

func makePanic() {
	defer func() { // !!!(4) добавим рековер через дефер, что бы вызвалась после паники
		panicValue := recover() // результатом функции реквер будет значение паники если паника произошла и Nil
		// если паника не произошла
		fmt.Println(panicValue)
	}()
	panic("какая то паника")
}

// Так же есть функция recover которая отлавливает нашу панику !!!(4)
