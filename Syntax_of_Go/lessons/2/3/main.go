// Ветвление (if, else, switch)

package main

import "fmt"

func main() {

	// Проверка на четность
	// for i := 0; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println("Четное")
	// 	} else {
	// 		fmt.Println("Нечетное")
	// 	}
	// }

	// проверка на деление без остатка
	// for i := 0; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Printf("Число %v делится нацело на 2 \n", i)
	// 	} else if i%3 == 0 {
	// 		fmt.Printf("Число %v делится нацело на 3 \n", i)
	// 	} else if i%5 == 0 {
	// 		fmt.Printf("Число %v делится нацело на 5 \n", i)
	// 	}
	// }

	// Выбор из вариантов
	for i := 0; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println("Ноль")
		case 1:
			fmt.Println("Один")
		case 2:
			fmt.Println("Два")
		case 3:
			fmt.Println("Три")
		case 4:
			fmt.Println("Четыре")
		case 5:
			fmt.Println("Пять")
		default:
			fmt.Println("Неизвестный номер")
		}
	}
}
