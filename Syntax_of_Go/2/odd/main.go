// 1. Написать функцию, которая определяет, четное ли число.

package main

import (
	"fmt"
)

//Функция проверяет на четность/нечетность принятый аргумент
func oddNumber() bool {
	var intNum int
	fmt.Print("Введите целое число: ")
	fmt.Scanln(&intNum)
	if intNum%2 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Эта программа проверяет введенное число на четность/нечетность")
	if oddNumber() == true {
		fmt.Println("Это число четное")
	} else {
		fmt.Println("Это число нечетное")
	}
}
