// Итерируемся по значениям массива

package main

import "fmt"

func main() {
	arr := [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	for i, item := range arr {
		fmt.Printf("%v - это %v элемент массива \n", item, i)
	}
}
