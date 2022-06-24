package chap7

import "fmt"

type SubjectScore struct {
	nationalLanguage int
	math             int
	english          int
	science          int
	society          int
}

type Score struct {
	Total   int
	Avarage float64
}

// SubjectScoreを引数にとってScoreを返す
func GoukeiToHeikin(args SubjectScore) Score {
	total := args.nationalLanguage + args.math + args.english + args.science + args.society

	avarage := float64(total) / 5

	return Score{total, avarage}
}

type Student struct {
	name   string
	result string
}

//名前と成績を受け取ったら文字列を返す
func Seiseki(s Student) string {
	return fmt.Sprintf("%sさんの評価は%sです。", s.name, s.result)
}

func (s *Student) Seiseki() string {
	return fmt.Sprintf("%sさんの評価は%sです。", s.name, s.result)
}

//座標
type Coordinate struct {
	x float64
	y float64
}

//x軸に対称な座標を返すメソッド
func (c *Coordinate) SymmetricX() Coordinate {
	c.y = c.y * -1
	return *c
}

// 1つの座標を引数にとって中点の座標を返す
func (c *Coordinate) Midpoint(c2 Coordinate) Coordinate {
	if c.x == c2.x && c.y == c2.y {
		x := c.x
		y := c.x
		return Coordinate{x: x, y: y}
	}

	x := (c.x + c2.x) / 2
	y := (c.y + c2.y) / 2

	return Coordinate{x: x, y: y}
}
