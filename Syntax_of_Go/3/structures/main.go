// 1. Описать несколько структур — любой легковой автомобиль и грузовик. Структуры должны
// содержать марку авто, год выпуска, объем багажника/кузова, запущен ли двигатель, открыты
// ли окна, насколько заполнен объем багажника.

// 2. Инициализировать несколько экземпляров структур. Применить к ним различные действия.
// Вывести значения свойств экземпляров в консоль.

package main

import (
	"Syntax_of_Go/3/structures/cars"
	"Syntax_of_Go/3/structures/lorrys"
	"fmt"
)

func main() {
	car1 := cars.Cars{"Fiat", 1966, 150, true, false, 110}
	car2 := cars.Cars{"BMW", 1995, 230, true, true, 50}
	lorry1 := lorrys.Lorrys{"Scania", 2007, 1500, false, true, 0}
	lorry2 := lorrys.Lorrys{"Mercedes", 2010, 2000, true, true, 1990}

	fmt.Println(car1)
	fmt.Println(car1.Model, "\tis engine start\t", car1.IsEngineStart)
	if car1.IsEngineStart == false {
		fmt.Println("Engine is stop")
	} else {
		fmt.Println("Wroom - wroom!!!")
	}
	car1.IsEngineStart = false
	if car1.IsEngineStart == false {
		fmt.Println("Engine is stop. End of driving")
	} else {
		fmt.Println("Wroom - wroom!!!")
	}
	fmt.Println(car1.Model, "\tis engine start\t", car1.IsEngineStart)
	fmt.Println(car2)
	fmt.Println(lorry1)
	fmt.Println(lorry2)

}
