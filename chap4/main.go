package main

import (
	"chap4/chap46"
	"fmt"
)

// 4-1から4-4
func baito_kyuyo(y int, t int) int {
	base := 850
	baseup := 100 * y
	base = base + baseup

	return base * t
}

func jikosyokai(n string) string {
	return fmt.Sprintf("Hello %s!", n)
}

func hyojyun_taiju(m float64) string {
	n := m * m * 22
	return fmt.Sprintf("%.2f", n)
}

func bmi(height float64, weight float64) string {
	b := weight / (height * height)

	return fmt.Sprintf("%.2f", b)
}

// ここまで

func main() {
	// 4-1から4-4
	year := 3
	work_time := 8 * 20

	fmt.Println(baito_kyuyo(year, work_time))

	name := "Tanaka"
	fmt.Println(jikosyokai(name))

	height := 1.85
	fmt.Println(hyojyun_taiju(height))

	weight := 78
	fmt.Println(bmi(height, float64(weight)))
	// ここまで

	//46
	fmt.Println(chap46.Tsuru_no_ashi(3))
	fmt.Println(chap46.Kame_no_ashi(4))
	fmt.Println(chap46.Tsurukame_no_ashi(3, 4))
	fmt.Println(chap46.Tsurukame(10, 36))

}
