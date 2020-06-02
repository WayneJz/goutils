package slice

import (
	"reflect"
	"testing"
)

func TestUniqueInt32s(t *testing.T) {
	type args struct {
		l []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "TestUniqueInt32s",
			args: args{
				l: []int32{0, 0, -1, 4, 7, 8, 9, 1, 4, 9, -4},
			},
			want: []int32{0, -1, 4, 7, 8, 9, 1, -4},
		},
		{
			name: "TestEmpty",
			args: args{
				l: []int32{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueInt32s(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueInt32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueInt64s(t *testing.T) {
	type args struct {
		l []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "TestUniqueInt64s",
			args: args{
				l: []int64{0, 0, -1, 4, 7, 8, 9, 1, 4, 9, -4},
			},
			want: []int64{0, -1, 4, 7, 8, 9, 1, -4},
		},
		{
			name: "TestEmpty",
			args: args{
				l: []int64{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueInt64s(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueInt64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueStrings(t *testing.T) {
	type args struct {
		l []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestUniqueStrings",
			args: args{
				l: []string{"", "if", "", "else", "else if", "else"},
			},
			want: []string{"", "if", "else", "else if"},
		},
		{
			name: "TestEmpty",
			args: args{
				l: []string{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueStrings(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueUint32s(t *testing.T) {
	type args struct {
		l []uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "TestUniqueUint32s",
			args: args{
				l: []uint32{0, 0, 4, 7, 8, 9, 1, 4, 9},
			},
			want: []uint32{0, 4, 7, 8, 9, 1},
		},
		{
			name: "TestEmpty",
			args: args{
				l: []uint32{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueUint32s(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueUint32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueUint64s(t *testing.T) {
	type args struct {
		l []uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "TestUniqueUint64s",
			args: args{
				l: []uint64{0, 0, 4, 7, 8, 9, 1, 4, 9},
			},
			want: []uint64{0, 4, 7, 8, 9, 1},
		},
		{
			name: "TestEmpty",
			args: args{
				l: []uint64{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueUint64s(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueUint64s() = %v, want %v", got, tt.want)
			}
		})
	}
}
