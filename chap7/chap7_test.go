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
