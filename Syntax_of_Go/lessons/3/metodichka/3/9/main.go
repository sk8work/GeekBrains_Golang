// Знакомимся с Json
// JSON (JavaScript Object Notation) — это простой формат хранения данных. Синтаксически он похож на
// объекты и листы в JavaScript. Часто он служит для передачи данных между сайтом и
// фронтенд-приложениями.
// Читаем и пишем Json
// Для записи данных в JSON служит метод Marshal из стандартной библиотеки. (пакет encoding/json)
// На вход он принимает любой объект, а возвращает массив из байт, а также ошибку.

package main

import (
	"Syntax_of_Go/lessons/3/metodichka/3/9/message"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	m := message.Message{"Alice", "Hello", 129470395881547000}

	b, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(b)

	fmt.Println(string(b))
}
