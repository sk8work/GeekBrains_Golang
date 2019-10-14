package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	jsonString := "{\"Name\": \"Alice\", \"Body\": \"Hello\", \"Time\": 129470395881547000}"

	b := []byte(jsonString)
	m := Message{}

	err := json.Unmarshal(b, &m)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(m)
}
