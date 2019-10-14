// Работа с функциями (func, defer, panic, recover)
// Возврат значений
package main

import "fmt"

func main() {
	ourslice := []int{55, 64, 98, 12, 35, 10}
	res := average(ourslice...)
	fmt.Println(res)
	res = average(12, 20, 50, 1, 5, 80)
	fmt.Println(res)
}

func average(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total / int(len(values))
}
