// Даны катеты прямоугольного треугольника. Найти его площадь, периметр и гипотенузу.
// Используйте тип данных float64 и функции из пакета math.

package main

import (
	"fmt"
	"math"
)

func main() {
	var catet1, catet2, hip, square, perimetr float64 = 0, 0, 0, 0, 0

	fmt.Print("Введите длину катета №1: ")
	fmt.Scan(&catet1)
	fmt.Print("Введите длину катета №2: ")
	fmt.Scan(&catet2)

	hip = math.Sqrt(math.Pow(catet1, 2) + math.Pow(catet2, 2))
	square = (catet1 * catet2) * 0.5
	perimetr = catet1 + catet2 + hip

	fmt.Println("Гипотенуза:", hip)
	fmt.Println("Площадь:", square)
	fmt.Println("Периметр:", perimetr)

}
