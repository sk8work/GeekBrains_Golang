package main

import (
	"fmt"

	"Syntax_of_Go/lessons/3/metodichka/3/4/stack"
)

func main() {
	stack.Push("Этот текст")
	stack.Push("Будет находиться в стеке")
	stack.Push("До первого возвращения к pop")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	stack.Push("Добавим еще текста")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
