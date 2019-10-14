// Обход по сущности

package main

import "fmt"

func main() {
	str := "Test string"
	for i, ch := range str {
		fmt.Println(i, "\t", string(ch), "\t", ch)
	}
	fmt.Println("")
	str2 := "Тестовая строка"
	for i, ch := range str2 {
		fmt.Println(i, "\t", string(ch), "\t", ch)
	}
}
