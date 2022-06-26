package chap81

import (
	"testing"
)

func TestPerson_t_String(t *testing.T) {
	type fields struct {
		Name       string
		Height     float64
		Weight     float64
		BirthMonth int
		BirthDay   int
		BloodType  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "問題83 テスト1",
			fields: fields{
				Name:       "mogi",
				Height:     180,
				Weight:     80,
				BirthMonth: 4,
				BirthDay:   1,
				BloodType:  "B",
			},
			want: "mogi,180.0cm,80.0kg,4月1日,B型",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Person_t{
				Name:       tt.fields.Name,
				Height:     tt.fields.Height,
				Weight:     tt.fields.Weight,
				BirthMonth: tt.fields.BirthMonth,
				BirthDay:   tt.fields.BirthDay,
				BloodType:  tt.fields.BloodType,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Person_t.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_t_Ketsueki_hyoji(t *testing.T) {
	type fields struct {
		Name       string
		Height     float64
		Weight     float64
		BirthMonth int
		BirthDay   int
		BloodType  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "問題84テスト1",
			fields: fields{
				Name:       "田中",
				Height:     180,
				Weight:     80,
				BirthMonth: 3,
				BirthDay:   1,
				BloodType:  "A",
			},
			want: "田中さんの血液型はA型です",
		},
		{
			name: "問題84テスト2",
			fields: fields{
				Name:       "佐藤",
				Height:     180,
				Weight:     80,
				BirthMonth: 3,
				BirthDay:   1,
				BloodType:  "B",
			},
			want: "佐藤さんの血液型はB型です",
		},
		{
			name: "問題84テスト3",
			fields: fields{
				Name:       "山本",
				Height:     180,
				Weight:     80,
				BirthMonth: 3,
				BirthDay:   1,
				BloodType:  "O",
			},
			want: "山本さんの血液型はO型です",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Person_t{
				Name:       tt.fields.Name,
				Height:     tt.fields.Height,
				Weight:     tt.fields.Weight,
				BirthMonth: tt.fields.BirthMonth,
				BirthDay:   tt.fields.BirthDay,
				BloodType:  tt.fields.BloodType,
			}
			if got := p.Ketsueki_hyoji(); got != tt.want {
				t.Errorf("Person_t.Ketsueki_hyoji() = %v, want %v", got, tt.want)
			}
		})
	}
}
