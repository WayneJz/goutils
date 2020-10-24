package condition

import (
	"testing"
)

func TestAny(t *testing.T) {
	type args struct {
		conditions []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestAny",
			args: args{
				conditions: []bool{
					1 == 1,
					false == true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.args.conditions...); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	type args struct {
		conditions []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestAll",
			args: args{
				conditions: []bool{
					1 == 1,
					false == true,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.args.conditions...); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
