// Программа для конвертирования рублей в доллары по фиксированному курсу 65.00 руб/USD

package main

import (
	"fmt"
)

func main() {
	const USDCourseDelimeter float32 = 65.00
	var rubs float32 = 0.00
	var dollars float32 = 0.00

	fmt.Print("Введите сумму в рублях, которую вы хотите обменять на доллары: ")
	fmt.Scan(&rubs)

	dollars = rubs / USDCourseDelimeter
	fmt.Println("Вы получили", dollars, "долларов")
}
