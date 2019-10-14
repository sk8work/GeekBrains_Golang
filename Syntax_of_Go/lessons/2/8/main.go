// Работа с функциями (func, defer, panic, recover)
// Паника и способ перестать паниковать (panic, recover)

package main

import "fmt"

func main() {
	f()
	fmt.Println("Выполнение f завершилось нормально")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Работа восстановлена f", r)
		}
	}()
	fmt.Println("Вызов g.")
	g(0)
	fmt.Println("Выполнение g завершено нормально.")
}

func g(i int) {
	if i > 5 {
		fmt.Println("Паника!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Отложенный вызов в g", i)
	fmt.Println("Вывод в g", i)
	g(i + 1)
}
