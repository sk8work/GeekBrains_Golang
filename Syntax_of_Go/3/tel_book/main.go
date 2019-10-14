// 4. * Внести в телефонный справочник дополнительные данные. Реализовать сохранение
// json-файла на диске с помощью пакета ioutil при изменении данных.

package main

import "Syntax_of_Go/3/tel_book/bookconstructor"

func main() {

	person1 := bookconstructor.Handbook{"Vasya", "Pupkin", 89545896547}
	person2 := bookconstructor.Handbook{"Grisha", "Veselov", 89581578965}
	person3 := bookconstructor.Handbook{"Vova", "Putin", 81354897854}
	person4 := bookconstructor.Handbook{"Dima", "Medvedev", 86593216597}
	person5 := bookconstructor.Handbook{"Igor", "Chaika", 87856289214}
	person6 := bookconstructor.Handbook{"Mariya", "Zakharova", 83965871587}
	person7 := bookconstructor.Handbook{"Vova", "Zhirinovsky", 86853742915}

}
