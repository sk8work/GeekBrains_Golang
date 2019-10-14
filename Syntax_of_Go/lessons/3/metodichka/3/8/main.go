package main

import "fmt"

type сircle struct {
	// структура Круг
	redius int
}

func main() {
	first := сircle{10}
	second := сircle{5}

	fmt.Println(first == second)
	fmt.Println(first != second)
}
