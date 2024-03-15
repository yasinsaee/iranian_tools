package util

import (
	"testing"
)

func TestConvertNumberToText(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "convert number to string",
			args:    args{number: "12۴"},
			want:    "صد و بیست و چهار",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertNumberToText(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertNumberToText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertNumberToText() = %v, want %v", got, tt.want)
			}
		})
	}
}
