package chap52

import (
	"errors"
)

//時間xを受け取ったら午前か午後かを返す
func Jikan(x int) (string, error) {
	if 0 <= x && x <= 12 {
		return "午前", nil
	} else if 13 <= x && x <= 23 {
		return "午後", nil
	} else {
		return "", errors.New("0から23までの数値を入力してください")
	}
}

// 誕生日を月mと日dを受け取ったら星座を返す
func Seiza(m, d int) (string, error) {
	if 0 >= m || 13 <= m {
		return "", errors.New("正しい日付を入力してください")
	}
	if 0 >= d || 32 <= d {
		return "", errors.New("正しい日付を入力してください")
	}

	switch m {
	case 3:
		if d <= 20 {
			return "うお座", nil
		} else {
			return "おひつじ座", nil
		}
	case 4:
		if d <= 19 {
			return "おひつじ座", nil
		} else if d <= 30 {
			return "おうし座", nil
		} else {
			return "", errors.New("正しい日付を入力してください")
		}
	case 5:
		if d <= 20 {
			return "おうし座", nil
		} else {
			return "ふたご座", nil
		}
	case 6:
		if d <= 21 {
			return "ふたご座", nil
		} else if d <= 30 {
			return "かに座", nil
		} else {
			return "", errors.New("正しい日付を入力してください")
		}
	case 7:
		if d <= 22 {
			return "かに座", nil
		} else {
			return "しし座", nil
		}
	case 8:
		if d <= 22 {
			return "しし座", nil
		} else {
			return "おとめ座", nil
		}
	case 9:
		if d <= 23 {
			return "おとめ座", nil
		} else if d <= 30 {
			return "てんびん座", nil
		} else {
			return "", errors.New("正しい日付を入力してください")
		}
	case 10:
		if d <= 23 {
			return "てんびん座", nil
		} else {
			return "さそり座", nil
		}
	case 11:
		if d <= 22 {
			return "さそり座", nil
		} else if d <= 30 {
			return "いて座", nil
		} else {
			return "", errors.New("正しい日付を入力してください")
		}
	case 12:
		if d <= 21 {
			return "いて座", nil
		} else {
			return "やぎ座", nil
		}
	case 1:
		if d <= 19 {
			return "やぎ座", nil
		} else {
			return "みずがめ座", nil
		}
	case 2:
		if d <= 20 {
			return "みずがめ座", nil
		} else {
			return "うお座", nil
		}
	}

	return "なし", errors.New("0から23までの数値を入力してください")
}
