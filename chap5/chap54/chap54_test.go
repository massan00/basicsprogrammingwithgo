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
