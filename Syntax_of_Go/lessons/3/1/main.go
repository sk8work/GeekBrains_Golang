// Массивы

package main

import "fmt"

func main() {
	var arr [10]int // Массив
	fmt.Println(len(arr), cap(arr))

	arr1 := [5]int{1, 2, 3, 4, 5} // Массив
	fmt.Printf("%T", arr1)
	fmt.Println()

	emptyarr := [5]int{}
	fmt.Println(emptyarr)

	sl := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", sl)
}
