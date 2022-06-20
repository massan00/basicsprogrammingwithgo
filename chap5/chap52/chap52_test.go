package chap52

import "testing"

func TestJikan(t *testing.T) {
	type args struct {
		time int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "11時",
			args: args{time: 11},
			want: "午前",
		},
		{
			name: "14時",
			args: args{time: 14},
			want: "午後",
		},
		{
			name: "20時",
			args: args{time: 20},
			want: "午後",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Jikan(tt.args.time); got != tt.want {
				t.Errorf("jikan = %s, want = %s", got, tt.want)
			}
		})
	}
}

func TestSeiza(t *testing.T) {
	type args struct {
		month int
		day   int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "4月20日",
			args: args{month: 4, day: 20},
			want: "おうし座",
		},
		{
			name: "8月30日",
			args: args{month: 8, day: 30},
			want: "おとめ座",
		},
		{
			name: "12月15日",
			args: args{month: 12, day: 15},
			want: "いて座",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Seiza(tt.args.month, tt.args.day); got != tt.want {
				t.Errorf("Seiza() = %s, want = %s", got, tt.want)
			}
		})
	}
}
