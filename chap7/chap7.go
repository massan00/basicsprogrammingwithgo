package chap7

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
