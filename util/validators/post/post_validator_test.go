package post

import "testing"

func TestPostalCodeValidator(t *testing.T) {
	type args struct {
		postalCode string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{postalCode: "5173985889"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PostalCodeValidator(tt.args.postalCode); got != tt.want {
				t.Errorf("PostalCodeValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}
