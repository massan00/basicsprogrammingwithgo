package chap54

import "testing"

func TestHanbetsushiki(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "問題54 test1",
			args: args{a: 2.0, b: 3.0, c: 4.0},
			want: -23.0,
		},
		{
			name: "問題54 test2",
			args: args{a: 1.0, b: 5.0, c: 4.0},
			want: 9.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hanbetsushiki(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("Hanbetsushiki() = %f, want= %f", got, tt.want)
			}
		})
	}

}

func TestKainokosuu(t *testing.T) {

	type args struct {
		a float64
		b float64
		c float64
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "問題55 test1",
			args: args{a: 2.0, b: 3.0, c: 4.0},
			want: 0,
		},
		{
			name: "問題55 test2",
			args: args{a: 1.0, b: 5.0, c: 4.0},
			want: 2,
		},
		{
			name: "問題55 test3",
			args: args{a: 1.0, b: 2.0, c: 4.0},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Kai_no_kosuu(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("Kai_no_kosuu() = %d, want= %d", got, tt.want)
			}
		})
	}
}

func TestKyosuukai(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "問題56 test1",
			args: args{a: 2.0, b: 3.0, c: 4.0},
			want: true,
		},
		{
			name: "問題56 test2",
			args: args{a: 2.0, b: -4.0, c: 2.0},
			want: false,
		},
		{
			name: "問題56 test3",
			args: args{a: 1.0, b: 2.0, c: 4.0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Kyosuukai(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("Kyosuukai() = %v, want= %v", got, tt.want)
			}
		})
	}

}

func TestTakei(t *testing.T) {
	type args struct {
		m  float64
		kg float64
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "身長1.75 体重 66.5",
			args: args{m: 1.75, kg: 66.5},
			want: "標準",
		},
		{
			name: "身長1.85 体重 60",
			args: args{m: 1.85, kg: 60},
			want: "やせ",
		},
		{
			name: "身長1.65 体重 80",
			args: args{m: 1.65, kg: 80},
			want: "肥満",
		},
		{
			name: "身長1.55 体重 90",
			args: args{m: 1.55, kg: 90},
			want: "高度肥満",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Taikei(tt.args.m, tt.args.kg); got != tt.want {
				t.Errorf("Taikei() = %s, want = %s", got, tt.want)
			}
		})
	}
}
