package main

import (
	"Syntax_of_Go/4/3/calculator"
	"fmt"
)

func main() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "help" {
			fmt.Println(calculator.Help())
			continue
		}

		if input == "exit" {
			break
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисления")
		}
	}
}
