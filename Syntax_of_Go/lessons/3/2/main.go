// слайсы

package main

import "fmt"

func main() {
	arr := make([]int, 0)
	arr1 := []int{1, 2, 3, 4, 5}
	fmt.Println(arr, cap(arr), len(arr))
	fmt.Println(arr1, cap(arr1), len(arr1))
}
