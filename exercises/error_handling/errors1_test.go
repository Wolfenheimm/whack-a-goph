package error_handling

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		wantErr bool
	}{
		{
			name:    "valid division",
			a:       10,
			b:       2,
			want:    5,
			wantErr: false,
		},
		{
			name:    "division by zero",
			a:       10,
			b:       0,
			want:    0,
			wantErr: true,
		},
		{
			name:    "negative numbers",
			a:       -10,
			b:       2,
			want:    -5,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
