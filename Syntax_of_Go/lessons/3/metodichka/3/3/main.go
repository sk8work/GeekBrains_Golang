// Срезы
package main

func main() {
	// Используем функцию make
	arr := make([]int, 5)
	//arr2 := make([]int, 5, 100)

	// Изменяем срезы: функции append, copy
	arr = append(arr, 6, 7, 8)

	// Создание нового среза на основе старого с дополнительными элементами
	//newArr := append(arr, 10, 10)
	var newArr []int
	copy(newArr, arr)

	// Изменение элементов среза производится с помощью индекса.
	arr[3] = 25
}
