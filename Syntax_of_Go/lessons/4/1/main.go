// Методы
// Методы напоминают функции, но они объявляются и работают вместе со структурами.
// Разбираем объявление методов

package main

import (
	"fmt"
	"math"
)

// Circle - простая структура для описания круга
type Circle struct {
	radius int
}

// Area - метод, котоный получает структуру по значению
// и возвращает ее площадь
func (c Circle) Area() float64 {
	return math.Pi * float64(c.radius) * float64(c.radius)
}

// ChangeRadius - это метод , который может изменить
// значение радиуса в конкретной структуре Circle
func (c *Circle) ChangeRadius(radius int) {
	c.radius = radius
}

func main() {
	circle1 := Circle{50}
	fmt.Println(circle1)
}
