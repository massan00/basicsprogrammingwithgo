package chap46

import (
	"testing"
)

func TestTsurunoashi(t *testing.T) {
	type args struct {
		tsuru int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "鶴4匹足8本",
			args: args{tsuru: 4},
			want: 8,
		},
		{
			name: "鶴8匹足16本",
			args: args{tsuru: 8},
			want: 16,
		},
		{
			name: "鶴13匹足26本",
			args: args{tsuru: 13},
			want: 26,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tsuru_no_ashi(tt.args.tsuru); got != tt.want {
				t.Errorf("tsuru_no_ashi = %d , want = %d", got, tt.want)
			}
		})
	}

}

func TestKamenoshi(t *testing.T) {
	type args struct {
		kame int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "亀3匹足12本",
			args: args{kame: 3},
			want: 12,
		},
		{
			name: "亀7匹",
			args: args{kame: 7},
			want: 28,
		},
		{
			name: "亀12匹",
			args: args{kame: 12},
			want: 48,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Kame_no_ashi(tt.args.kame); got != tt.want {
				t.Errorf("kame_no_ashi = %d, want = %d", got, tt.want)
			}
		})
	}
}

func TestTsurukamenoashi(t *testing.T) {
	type args struct {
		tsuru int
		kame  int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "鶴4匹亀3匹",
			args: args{tsuru: 4, kame: 3},
			want: 20,
		},
		{
			name: "鶴8匹亀11匹",
			args: args{tsuru: 8, kame: 11},
			want: 60,
		},
		{
			name: "鶴0匹亀1匹",
			args: args{tsuru: 0, kame: 1},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tsurukame_no_ashi(tt.args.tsuru, tt.args.kame); got != tt.want {
				t.Errorf("turukamonoashi = %d, want= %d", got, tt.want)
			}
		})
	}

}

func TestTsurukame(t *testing.T) {
	type args struct {
		tsurukame int
		footsum   int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "鶴2亀4",
			args: args{tsurukame: 6, footsum: 20},
			want: 2,
		},
		{
			name: "鶴1亀1",
			args: args{tsurukame: 2, footsum: 6},
			want: 1,
		},
		{
			name: "鶴3亀2",
			args: args{tsurukame: 5, footsum: 14},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tsurukame(tt.args.tsurukame, tt.args.footsum); got != tt.want {
				t.Errorf("tsurukame() = %d, want= %d", got, tt.want)
			}
		})
	}
}
