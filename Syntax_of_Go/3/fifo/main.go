// 3. * Реализовать очередь. Это структура данных, работающая по принципу FIFO (First Input —
// 	first output, или «первым зашел — первым вышел»).

package main

import (
	"Syntax_of_Go/3/fifo/fifo"
	"fmt"
)

func main() {
	fifo.Push("Первый")
	fifo.Push("Второй")
	fifo.Push("Третий")

	fmt.Println(fifo.Shift())
	fmt.Println(fifo.Shift())

	fifo.Push("Четвертый")

	fmt.Println(fifo.Shift())
	fmt.Println(fifo.Shift())
	fmt.Println(fifo.Shift())

	fifo.Push("Пятый")
	fifo.Push("Шестой")

	fmt.Println(fifo.Shift())
	fmt.Println(fifo.Shift())
	fmt.Println(fifo.Shift())
}
