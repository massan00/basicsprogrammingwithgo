package chap54

//2次方程式ax^2+bx+c=0の係数a,b,c(全て実数)を与えられたら判別式の値を返す
func Hanbetsushiki(a, b, c float64) float64 {
	return (b * b) - (4 * a * c)
}

//2次方程式ax^2+bx+c=0の係数a,b,c(全て実数)を与えられたら解の個数を返す
func Kai_no_kosuu(a, b, c float64) int {
	num := Hanbetsushiki(a, b, c)
	if num > 0 {
		return 2
	} else if num == 0 {
		return 1
	} else {
		return 0
	}
}

//2次方程式ax^2+bx+c=0の係数a,b,c(全て実数)を与えられたら虚数解を持つかどうかを判定する
func Kyosuukai(a, b, c float64) bool {
	if Hanbetsushiki(a, b, c) < 0 {
		return true
	} else {
		return false
	}
}
