package main

import (
	"chap8/chap81"
	"fmt"
	"time"
)

func main() {
	b1 := chap81.Book_t{
		Title:     "プログラミングの基礎",
		Author:    "浅井健一",
		Publisher: "サイエンス社",
		Price:     2300,
		Isbn:      9784781911601,
	}
	fmt.Println(b1)

	o1 := chap81.Okozukai_t{
		Item:  "本",
		Price: 2300,
		Place: "東京",
		Date:  time.Date(2022, 6, 22, 0, 0, 0, 0, time.Local),
	}
	fmt.Println(o1)

	p1 := chap81.Person_t{
		Height:     180,
		Weight:     80,
		BirthMonth: 1,
		BirthDay:   1,
		BloodType:  "A",
	}
	fmt.Println(p1)
}
