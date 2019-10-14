// * Пользователь вводит сумму вклада в банк и годовой процент.
// Найти сумму вклада через 5 лет.

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Вы получили сумму:", calcDeposit())
}

func calcDeposit() float64 {
	var startSum float64
	var years float64
	var yearPercent float64
	var result float64
	fmt.Print("Введите сумму в рублях, которую вы хотите положить на депозит: ")
	fmt.Scan(&startSum)
	fmt.Print("Введите срок вклада (в годах): ")
	fmt.Scan(&years)
	fmt.Print("Введите процент по вкладу: ")
	fmt.Scan(&yearPercent)
	result = startSum * math.Pow((1+yearPercent/100), years)
	return result
}
