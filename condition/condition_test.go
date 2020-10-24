package condition

import (
	"reflect"
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
			name: "TestAny 1",
			args: args{
				conditions: []bool{
					1 == 1,
					false == true,
				},
			},
			want: true,
		},
		{
			name: "TestAny 2",
			args: args{
				conditions: []bool{
					1 == 2,
				},
			},
			want: false,
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
			name: "TestAll 1",
			args: args{
				conditions: []bool{
					1 == 1,
					false == true,
				},
			},
			want: false,
		},
		{
			name: "TestAll 2",
			args: args{
				conditions: []bool{
					1 == 1,
					true == true,
				},
			},
			want: true,
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

func TestTernary(t *testing.T) {
	type args struct {
		condition bool
		a         interface{}
		b         interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "TestTernary 1",
			args: args{
				condition: 1 == 1,
				a:         "haha",
				b:         1,
			},
			want: "haha",
		},
		{
			name: "TestTernary 2",
			args: args{
				condition: 1 == 2,
				a:         "haha",
				b:         1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ternary(tt.args.condition, tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ternary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuaternary(t *testing.T) {
	type args struct {
		condition1 bool
		condition2 bool
		both       interface{}
		cond1      interface{}
		cond2      interface{}
		neither    interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "TestQuaternary 1",
			args: args{
				condition1: 1 == 1,
				condition2: 2 == 2,
				both:       "both",
				cond1:      "cond1",
				cond2:      "cond2",
				neither:    "neither",
			},
			want: "both",
		},
		{
			name: "TestQuaternary 2",
			args: args{
				condition1: 1 == 1,
				condition2: 2 != 2,
				both:       "both",
				cond1:      "cond1",
				cond2:      "cond2",
				neither:    "neither",
			},
			want: "cond1",
		},
		{
			name: "TestQuaternary 3",
			args: args{
				condition1: 1 != 1,
				condition2: 2 == 2,
				both:       "both",
				cond1:      "cond1",
				cond2:      "cond2",
				neither:    "neither",
			},
			want: "cond2",
		},
		{
			name: "TestQuaternary 3",
			args: args{
				condition1: 1 != 1,
				condition2: 2 != 2,
				both:       "both",
				cond1:      "cond1",
				cond2:      "cond2",
				neither:    "neither",
			},
			want: "neither",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Quaternary(tt.args.condition1, tt.args.condition2, tt.args.both, tt.args.cond1, tt.args.cond2, tt.args.neither); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quaternary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasZeroValue(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestHasZeroValue 1",
			args: args{
				values: []interface{}{nil},
			},
			want: true,
		},
		{
			name: "TestHasZeroValue 2",
			args: args{
				values: []interface{}{0},
			},
			want: true,
		},
		{
			name: "TestHasZeroValue 3",
			args: args{
				values: []interface{}{[]int64{}},
			},
			want: false,
		},
		{
			name: "TestHasZeroValue 4",
			args: args{
				values: []interface{}{func() *int {
					a := 0
					return &a
				}},
			},
			want: false,
		},
		{
			name: "TestHasZeroValue 5",
			args: args{
				values: []interface{}{""},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasZeroValue(tt.args.values...); got != tt.want {
				t.Errorf("HasZeroValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestHasNumber 1",
			args: args{
				s: " ew1bq=#!>",
			},
			want: true,
		},
		{
			name: "TestHasNumber 2",
			args: args{
				s: "\u2345",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasNumber(tt.args.s); got != tt.want {
				t.Errorf("HasNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasLetter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestHasLetter 1",
			args: args{
				s: " !@#>'$123",
			},
			want: false,
		},
		{
			name: "TestHasLetter 2",
			args: args{
				s: "\u2345",
			},
			want: false,
		},
		{
			name: "TestHasLetter 3",
			args: args{
				s: "e\u2345",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasLetter(tt.args.s); got != tt.want {
				t.Errorf("HasLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
