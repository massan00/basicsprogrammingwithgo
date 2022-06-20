package main

import (
	"chap5/chap52"
	"fmt"
	"log"
)

func main() {
	// 練習問題5.2
	h, err := chap52.Jikan(20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(h)

	// 練習問題5.3
	s, err := chap52.Seiza(6, 20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
