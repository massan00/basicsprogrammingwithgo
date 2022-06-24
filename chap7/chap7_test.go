package chap7

import "testing"

func TestGoukeiToHeikin(t *testing.T) {
	tests := []struct {
		name string
		args SubjectScore
		want Score
	}{
		{
			name: "問題71 テスト1",
			args: SubjectScore{
				nationalLanguage: 80,
				math:             80,
				english:          80,
				science:          80,
				society:          80,
			},
			want: Score{Total: 400, Avarage: 80},
		},
		{
			name: "問題71 テスト2",
			args: SubjectScore{
				nationalLanguage: 90,
				math:             60,
				english:          55,
				science:          88,
				society:          76,
			},
			want: Score{Total: 369, Avarage: 73.8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GoukeiToHeikin(tt.args); got != tt.want {
				t.Errorf("GoukeiToHeikin() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestSeiseki(t *testing.T) {
	tests := []struct {
		name string
		args Student
		want string
	}{
		{
			name: "問題72 テスト1",
			args: Student{name: "sato", result: "A"},
			want: "satoさんの評価はAです。",
		},
		{
			name: "問題72 テスト2",
			args: Student{name: "tanaka", result: "B"},
			want: "tanakaさんの評価はBです。",
		},
		{
			name: "問題72 テスト3",
			args: Student{name: "sasaki", result: "C"},
			want: "sasakiさんの評価はCです。",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Seiseki(tt.args); got != tt.want {
				t.Errorf("Seiseki()=%s, want = %s", got, tt.want)
			}
		})
	}
}

func TestSeisekim(t *testing.T) {
	tests := []struct {
		name string
		args Student
		want string
	}{
		{
			name: "メソッド 問題72 テスト1",
			args: Student{name: "sato", result: "A"},
			want: "satoさんの評価はAです。",
		},
		{
			name: "メソッド 問題72 テスト2",
			args: Student{name: "tanaka", result: "B"},
			want: "tanakaさんの評価はBです。",
		},
		{
			name: "メソッド 問題72 テスト3",
			args: Student{name: "sasaki", result: "C"},
			want: "sasakiさんの評価はCです。",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.Seiseki(); got != tt.want {
				t.Errorf("Seiseki()=%s, want = %s", got, tt.want)
			}
		})
	}
}

func TestSymmetricX(t *testing.T) {
	tests := []struct {
		name string
		args Coordinate
		want Coordinate
	}{
		{
			name: "問題73 テスト1",
			args: Coordinate{x: 1, y: 1},
			want: Coordinate{x: 1, y: -1},
		},
		{
			name: "問題73 テスト2",
			args: Coordinate{x: 0, y: 0},
			want: Coordinate{x: 0, y: 0},
		},
		{
			name: "問題73 テスト3",
			args: Coordinate{x: 10, y: -5},
			want: Coordinate{x: 10, y: 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.SymmetricX(); got != tt.want {
				t.Errorf("SymmetricX = %v, want = %v", got, tt.want)
			}
		})
	}

}

func TestMidpoint(t *testing.T) {
	tests := []struct {
		name string
		obj  Coordinate
		args Coordinate
		want Coordinate
	}{
		{
			name: "問題74 テスト1",
			obj:  Coordinate{x: 1, y: 1},
			args: Coordinate{x: 1, y: 1},
			want: Coordinate{x: 1, y: 1},
		},
		{
			name: "問題74 テスト2",
			obj:  Coordinate{x: 3, y: 3},
			args: Coordinate{x: -3, y: -3},
			want: Coordinate{x: 0, y: 0},
		},
		{
			name: "問題74 テスト3",
			obj:  Coordinate{x: 6, y: 4},
			args: Coordinate{x: 2, y: 0},
			want: Coordinate{x: 4, y: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.obj.Midpoint(tt.args); got != tt.want {
				t.Errorf("SymmetricX = %v, want = %v", got, tt.want)
			}
		})
	}

}
