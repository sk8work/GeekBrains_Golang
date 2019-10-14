// Мапы

package main

import (
	"fmt"
	"hash/crc32"
)

func hash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func findPosition(key string, length int) int {
	h := hash(key)
	return int(h % uint32(length))
}

func main() {
	arr := make([]int, 4)
	key := "some_key123"
	value := 42
	position := findPosition(key, len(arr))
	fmt.Println(position)
	fmt.Println(arr)
	arr[position] = value
	fmt.Println(arr)
}
