// Создать тип, описывающий контакт в телефонной книге. Создать псевдоним типа телефонной
// книги (массив контактов) и реализовать для него интерфейс Sort{}.

package main

import (
	"fmt"
	"sort"
)

type PersonContact struct {
	name  string
	age   int
	tel   int64
	email string
}

func (p PersonContact) String() string {
	return fmt.Sprintf("%s: %d, %s", p.name, p.tel, p.email)
}

type ByTel []PersonContact

func (a ByTel) Len() int           { return len(a) }
func (a ByTel) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTel) Less(i, j int) bool { return a[i].tel > a[j].tel }

func main() {
	people := []PersonContact{
		{"Mikky", 31, 111111, "qwe@kjh.ru"},
		{"Billy", 42, 1111112, "123@qwe.ru"},
		{"Villy", 17, 1111114, "err@her.ru"},
		{"Dilli", 26, 1111113, "err@her.ru"},
	}

	fmt.Println(people)
	sort.Sort(ByTel(people))
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].tel < people[j].tel
	})
	fmt.Println(people)
}
