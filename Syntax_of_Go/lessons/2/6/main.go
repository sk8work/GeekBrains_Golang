// Работа с функциями (func, defer, panic, recover)
// Функции могут возвращать несколько значений:

package main

import "fmt"

func main() {
	ourSlice := []int{10, 50, 30, 45, 32}

	// Получаем все аргументы
	res, count := average(10, 50, 20, 12, 4)
	fmt.Println(res)
	fmt.Println(count)

	// Игнорируем все аргументы
	average(ourSlice...)

	// Игнорируем только второй аргумент
	res, _ = average(10, 5, 2, 12, 3)
	fmt.Println(res)
	fmt.Println(count)
}

func average(values ...int) (res int, count int) {
	count = len(values)
	total := 0
	for _, v := range values {
		total += v
	}
	res = total / count
	return
}
