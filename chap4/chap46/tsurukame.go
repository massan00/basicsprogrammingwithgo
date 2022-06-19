package chap46

//鶴の数tsuruが与えられたら足の本数yを返す
func Tsuru_no_ashi(tsuru int) int {
	return tsuru * 2
}

//亀の数kameが与えられたら足の本数yを返す
func Kame_no_ashi(kame int) int {
	return kame * 4
}

// 鶴の数tsuruと亀の数kameが与えられたら足の数の合計を返す
func Tsurukame_no_ashi(tsuru, kame int) int {
	return tsuru*2 + kame*4
}

//鶴と亀の合計xと足の合計yが与えられたら鶴の数を返す
func Tsurukame(x, y int) int {
	tsuru := (x * 2) - (y / 2)
	return tsuru
}
