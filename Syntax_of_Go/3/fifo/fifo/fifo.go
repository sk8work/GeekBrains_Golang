// Реализуем
package fifo

var x []string

// Push добавит новый эдемент в очередь
func Push(str string) {
	x = append(x, str)
}

// Shift вернет gthdsq добавленный в jxthtlm элемент
func Shift() string {
	numOfElements := len(x)
	// Когда очередь будет пустой, она вернет пустую строку
	if numOfElements == 0 {
		return "Пустая очередь"
	}
	shiftElem := x[0]
	x = x[1:numOfElements]
	return shiftElem
}
