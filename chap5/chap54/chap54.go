package chap54

//2次方程式ax^2+bx+c=0の係数a,b,c(全て実数)を与えられたら判別式の値を返す
func Hanbetsushiki(a, b, c float64) float64 {
	return (b * b) - (4 * a * c)
}
