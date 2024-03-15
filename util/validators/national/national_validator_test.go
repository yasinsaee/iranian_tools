package national

import "testing"

func TestNationalCodeValidator_Validate(t *testing.T) {
	type fields struct {
		NationalCode string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "",
			fields: fields{NationalCode: "1234567891"},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &NationalCodeValidator{
				NationalCode: tt.fields.NationalCode,
			}
			if got := v.Validate(); got != tt.want {
				t.Errorf("NationalCodeValidator.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
