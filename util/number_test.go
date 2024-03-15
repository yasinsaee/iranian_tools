package util

import "testing"

func TestChangeToPersianDigit(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test change to persian digit",
			args: args{number: "123456"},
			want: "۱۲۳۴۵۶",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChangeToPersianDigit(tt.args.number); got != tt.want {
				t.Errorf("ChangeToPersianDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
