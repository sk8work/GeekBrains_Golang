// Написать функцию, которая последовательно выводит на экран чисkf Фибоначчи,
// начиная от 0.
// a. Числа Фибоначчи определяются соотношениями ​ Fn=Fn-1 + Fn-2.

package main

import "fmt"

func printFib(n int) {
	var fib1 uint64 = 0
	var fib2 uint64 = 1
	for i := 0; i < n; i++ {
		fmt.Println(fib1)
		fib1, fib2 = fib2, fib1+fib2
	}
}

func main() {
	fmt.Println("-> Эта программа выводит числа Фибоначчи")
	printFib(100)
	fmt.Println("-> Конец программы")
}
