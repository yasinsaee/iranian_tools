package phone

import "testing"

func TestCellphoneValidator(t *testing.T) {
	type args struct {
		cellphone string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{cellphone: "09039686577"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CellphoneValidator(tt.args.cellphone); got != tt.want {
				t.Errorf("CellphoneValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}
