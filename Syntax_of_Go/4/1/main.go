// Написать свой интерфейс и создать несколько структур
// удовлетворяющих ему

package main

import "fmt"

type Animal interface {
	say()
	play()
}

type Cat struct{}
type Dog struct{}

func (c Cat) say() {
	fmt.Println("Meow")
}
func (c Cat) play() {
	fmt.Println("Play with mouse toy")
}
func (d Dog) say() {
	fmt.Println("Woof-Woof")
}
func (d Dog) play() {
	fmt.Println("Play with bone")
}

func main() {
	var cat1 Animal = Cat{}
	var dog1 Animal = Dog{}
	cat1.say()
	cat1.play()
	dog1.say()
	dog1.play()
}
