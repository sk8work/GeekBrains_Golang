// Создаем и изменяем массив

package main

import "fmt"

func main() {
	// Объявление массива со щеачениями
	arr := [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	// Оббъявление пустого массива
	var arr2 [5]int

	arr2[2] = 10

	fmt.Println(arr)
	fmt.Println(arr2)
}
