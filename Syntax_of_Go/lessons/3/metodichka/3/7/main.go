// Структуры

package main

import "fmt"

type position struct {
	X int
	Y int
}

func main() {
	var FirstPoint = position{2, 4}
	var SecondPoint = position{Y: 3, X: 3}
	var ThirdPoint = position{
		Y: 6,
		X: 3,
	}
	fmt.Println(FirstPoint)
	fmt.Println(SecondPoint)
	fmt.Println(ThirdPoint)
}
