// 2. Написать функцию, которая определяет, делится ли число без остатка на 3.

package main

import (
	"fmt"
)

//Функция проверяет делится ли число на 3 без остатка
func oddNumber() bool {
	var intNum int
	fmt.Print("Введите целое число: ")
	fmt.Scanln(&intNum)
	if intNum%3 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Эта программа проверяет делится ли число на 3 без остатка")
	if oddNumber() == true {
		fmt.Println("Это число делится на 3 без остатка")
	} else {
		fmt.Println("Это число не делится на 3 без остатка")
	}
}
