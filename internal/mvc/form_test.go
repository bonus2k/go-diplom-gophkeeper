package mvc

import (
	"testing"
)

func Test_formatCardNumber(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "16 number characters",
			args: args{
				text: "1234123412341234",
			},
			want: "1234-1234-1234-1234",
		},
		{
			name: "15 number characters",
			args: args{
				text: "123412341234123",
			},
			want: "1234-1234-1234-123",
		},
		{
			name: "15 number characters with alphabetic characters",
			args: args{
				text: "D1F2AS341234SS123F41G2G3H",
			},
			want: "D1F2-AS34-1234-SS12-3F41-G2G3-H",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatCardNumber(tt.args.text); got != tt.want {
				t.Errorf("formatCardNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_form_note(t *testing.T) {

}
