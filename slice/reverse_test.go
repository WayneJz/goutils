package slice

import (
	"reflect"
	"testing"
)

func TestReverseStrings(t *testing.T) {
	type args struct {
		l []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestReverseStrings 1 odd length",
			args: args{
				l: []string{"one", "two", "three"},
			},
			want: []string{"three", "two", "one"},
		},
		{
			name: "TestReverseStrings 2 even length",
			args: args{
				l: []string{"one", "two", "three", "four"},
			},
			want: []string{"four", "three", "two", "one"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.args.l
			ReverseStrings(l)
			if !reflect.DeepEqual(l, tt.want) {
				t.Errorf("ReverseStrings() = %v, want %v", l, tt.want)
			}
		})
	}
}

func TestReverseInt64s(t *testing.T) {
	type args struct {
		l []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "TestReverseInt64s 1 odd length",
			args: args{
				l: []int64{1, 2, 3},
			},
			want: []int64{3, 2, 1},
		},
		{
			name: "TestReverseInt64s 2 even length",
			args: args{
				l: []int64{1, 2, 3, 4},
			},
			want: []int64{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.args.l
			ReverseInt64s(l)
			if !reflect.DeepEqual(l, tt.want) {
				t.Errorf("ReverseInt64s() = %v, want %v", l, tt.want)
			}
		})
	}
}

func TestReverseInt32s(t *testing.T) {
	type args struct {
		l []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "TestReverseInt32s 1 odd length",
			args: args{
				l: []int32{1, 2, 3},
			},
			want: []int32{3, 2, 1},
		},
		{
			name: "TestReverseInt32s 2 even length",
			args: args{
				l: []int32{1, 2, 3, 4},
			},
			want: []int32{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.args.l
			ReverseInt32s(l)
			if !reflect.DeepEqual(l, tt.want) {
				t.Errorf("ReverseInt32s() = %v, want %v", l, tt.want)
			}
		})
	}
}

func TestReverseUint64s(t *testing.T) {
	type args struct {
		l []uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "TestReverseUint64s 1 odd length",
			args: args{
				l: []uint64{1, 2, 3},
			},
			want: []uint64{3, 2, 1},
		},
		{
			name: "TestReverseUint64s 2 even length",
			args: args{
				l: []uint64{1, 2, 3, 4},
			},
			want: []uint64{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.args.l
			ReverseUint64s(l)
			if !reflect.DeepEqual(l, tt.want) {
				t.Errorf("ReverseUint64s() = %v, want %v", l, tt.want)
			}
		})
	}
}

func TestReverseUint32s(t *testing.T) {
	type args struct {
		l []uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "TestReverseUint32s 1 odd length",
			args: args{
				l: []uint32{1, 2, 3},
			},
			want: []uint32{3, 2, 1},
		},
		{
			name: "TestReverseUint32s 2 even length",
			args: args{
				l: []uint32{1, 2, 3, 4},
			},
			want: []uint32{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.args.l
			ReverseUint32s(l)
			if !reflect.DeepEqual(l, tt.want) {
				t.Errorf("ReverseUint32s() = %v, want %v", l, tt.want)
			}
		})
	}
}
