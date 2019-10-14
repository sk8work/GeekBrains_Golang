// Пишем простой калькулятор

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstNum float32
	var secondNum float32
	var action string

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	// Ввод первого числа
	firstNum = enterNumber("Введите первое число: ")
	action = inputData("Укажите действие: ")
	// Ввод второго числа
	secondNum = enterNumber("Введите второе число: ")
	fmt.Printf("Результат: %v \n", calculate(firstNum, secondNum, action))
}

func calculate(firstNum float32, secondNum float32, action string) float32 {
	switch action {
	case "+":
		return firstNum + secondNum
	case "-":
		return firstNum - secondNum
	case "*":
		return firstNum * secondNum
	case "/":
		return firstNum / secondNum
	default:
		fmt.Println("Не удалось распознать действие")
		return 0
	}
}

func inputData(msg string) (data string) {
	fmt.Print(msg)
	fmt.Scanln(&data)
	if data == "exit" {
		panic(nil)
	}
	return data
}

func enterNumber(msg string) float32 {
	for {
		num, err := strconv.ParseFloat(inputData(msg), 32)
		if err == nil {
			//Ввод числа прошел без ошибок
			return float32(num)
			break
		}
		fmt.Println("Не удалось распознать число")
	}
	return 0
}
