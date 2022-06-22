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

// bmi heightとweightを引数にとりbmi計算後数値を返す
func bmi(height float64, weight float64) float64 {
	b := weight / (height * height)

	return b
}

//身長mと体重kgを引数にとり、数値によって肥満状態を判定して文字列を返す
func Taikei(m float64, kg float64) string {
	bmi := bmi(m, kg)
	switch {
	case bmi < 18.5:
		return "やせ"
	case bmi >= 18.5 && bmi < 25:
		return "標準"
	case bmi >= 25 && bmi < 30:
		return "肥満"
	default:
		return "高度肥満"
	}
}
