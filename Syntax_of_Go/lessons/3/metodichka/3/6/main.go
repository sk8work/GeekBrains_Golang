// Комбинируем, делаем вложенные структуры, реализуем
//телефонный справочник

package main

import "fmt"

func main() {
	addressBook := make(map[string][]int)

	addressBook["Alex"] = []int{89996543210}
	addressBook["Bob"] = []int{89167243812}
	addressBook["Bob"] = append(addressBook["Bob"], 8915524327)

	fmt.Println(addressBook)

	for name, numbers := range addressBook {
		fmt.Println("Абонент:", name)
		for i, number := range numbers {
			fmt.Printf("\t %v) %v \n", i+1, number)
		}
	}
}
