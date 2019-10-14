// Работа с функциями (func, defer, panic, recover)
// Отложенный вызов (defer)

package main

import "fmt"

func average(values ...int) (res, count int) {
	count = len(values)
	total := 0
	for _, v := range values {
		total += v
	}
	res = total / count
	return
}

func main() {
	// // Будет вызвана последней
	// defer fmt.Println("First")
	// // Будет вызвана четвертой
	// defer fmt.Println("Second")
	// // Будет вызвана третьей
	// defer fmt.Println("Third")
	// // Будет вызвана первой
	// fmt.Println("Forth")
	// // Будет вызвана второй
	// fmt.Println("Fifth")

	// Второй пример
	// Будет вызвана последней
	defer fmt.Println(average(1, 2, 3))
	// Будет вызвана третьей
	defer fmt.Println(average(10, 20, 30))
	// Будет вызвана первой
	fmt.Println(average(1))
	// Будет вызвана второй
	fmt.Println(average(15, 23))
}
