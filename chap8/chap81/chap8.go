package chap81

import (
	"fmt"
	"time"
)

//問題81タイトル・著者名・出版社・値段・ISBNをもつ本のデータ
type Book_t struct {
	Title     string
	Author    string
	Publisher string
	Price     int
	Isbn      int
}

func (b Book_t) String() string {
	return fmt.Sprintf("%s,%s,%s,%d,%d", b.Title, b.Author, b.Publisher, b.Price, b.Isbn)
}

//問題82 お小遣い帳 買ったもの、値段、場所、日付
type Okozukai_t struct {
	Item  string
	Price int
	Place string
	Date  time.Time
}

func (o Okozukai_t) String() string {
	date := o.Date.Format("2006-01-02")
	return fmt.Sprintf("%s,%d,%s,%s", o.Item, o.Price, o.Place, date)
}

//問題83 身長、体重、誕生日、血液型を持つperson_t
type Person_t struct {
	Name       string
	Height     float64
	Weight     float64
	BirthMonth int
	BirthDay   int
	BloodType  string
}

func (p Person_t) String() string {
	return fmt.Sprintf("%s,%.1fcm,%.1fkg,%d月%d日,%s型", p.Name, p.Height, p.Weight, p.BirthMonth, p.BirthDay, p.BloodType)
}

func (p Person_t) Ketsueki_hyoji() string {
	return fmt.Sprintf("%sさんの血液型は%s型です", p.Name, p.BloodType)
}
