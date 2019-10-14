// Карты

package main

import "fmt"

func main() {
	// Создаем и изменяем карты, функция delete
	addressBook := make(map[string]int)
	// Чтобы добавить элемент в карту, нужно поступить следующим образом
	addressBook["Alex"] = 89991234567
	// Чтобы изменить элемент
	addressBook["Alex"] = 85551233213
	// Удаление элементов
	//delete(addressBook, "Alex")

	if number, ok := addressBook["Alex"]; ok {
		fmt.Println(number)
	}

	// Итерируемся по значениям
	for name, number := range addressBook {
		fmt.Println("Номер абонента", name, number)
	}
}
