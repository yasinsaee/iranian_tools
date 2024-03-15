package mail

import "testing"

func TestEmailValidator(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test email",
			args: args{email: "yasinvsaee@gmail.com"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EmailValidator(tt.args.email); got != tt.want {
				t.Errorf("EmailValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}
