package main

import (
	"chap9/chap91"
	"fmt"
)

func main() {
	//問題9.1
	season := [...]string{"春", "夏", "秋", "冬"}
	fmt.Println(season)
	//問題9.2
	persons := make([]chap91.Person_t, 3)
	persons[0] = chap91.Person_t{
		Name:       "satou",
		Height:     180,
		Weight:     80,
		BirthMonth: 4,
		BirthDay:   1,
		Blood:      "A",
	}

	persons[1] = chap91.Person_t{
		Name:       "katou",
		Height:     188,
		Weight:     90,
		BirthMonth: 5,
		BirthDay:   1,
		Blood:      "AB",
	}

	persons[2] = chap91.Person_t{
		Name:       "mutou",
		Height:     170,
		Weight:     65,
		BirthMonth: 6,
		BirthDay:   2,
		Blood:      "O",
	}

	fmt.Println(persons)
	fmt.Println(persons[0].Name)
}
