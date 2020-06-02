package slice

import (
	"testing"
)

func TestInStrings(t *testing.T) {
	type args struct {
		l []string
		s []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestInStrings 1",
			args: args{
				l: []string{"temp", "t", "", "empty", ""},
				s: []string{"temp", "", ""},
			},
			want: true,
		},
		{
			name: "TestInStrings 2",
			args: args{
				l: []string{"temp", "t", "", "empty", ""},
				s: []string{"temp", "", "tt"},
			},
			want: false,
		},
		{
			name: "TestInStrings s nil",
			args: args{
				l: []string{"temp", "t", "", "empty", ""},
				s: nil,
			},
			want: true,
		},
		{
			name: "TestInStrings both nil",
			args: args{
				l: nil,
				s: nil,
			},
			want: true,
		},
		{
			name: "TestInStrings l nil",
			args: args{
				l: nil,
				s: []string{"temp", "", ""},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InStrings(tt.args.l, tt.args.s...); got != tt.want {
				t.Errorf("InStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInInt64s(t *testing.T) {
	type args struct {
		l    []int64
		nums []int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestInInt64s 1",
			args: args{
				l:    []int64{0, 0, -1, 4, 7, -4},
				nums: []int64{0, 0, 0, 4, -4, -1, -1, -4},
			},
			want: true,
		},
		{
			name: "TestInInt64s 2",
			args: args{
				l:    []int64{0, 0, -1, 4, 7, -4},
				nums: []int64{0, 0, 0, 4, 4, -1, -1, -4, -7},
			},
			want: false,
		},
		{
			name: "TestInInt64s s nil",
			args: args{
				l:    []int64{0, 0, -1, 4, 7, -4},
				nums: nil,
			},
			want: true,
		},
		{
			name: "TestInInt64s both nil",
			args: args{
				l:    nil,
				nums: nil,
			},
			want: true,
		},
		{
			name: "TestInInt64s nums nil",
			args: args{
				l:    nil,
				nums: []int64{0, 0, 0, 4, 4, -1, -1, -4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InInt64s(tt.args.l, tt.args.nums...); got != tt.want {
				t.Errorf("InInt64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInInt32s(t *testing.T) {
	type args struct {
		l    []int32
		nums []int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestInInt32s 1",
			args: args{
				l:    []int32{0, 0, -1, 4, 7, -4},
				nums: []int32{0, 0, 0, 4, -4, -1, -1, -4},
			},
			want: true,
		},
		{
			name: "TestInInt32s 2",
			args: args{
				l:    []int32{0, 0, -1, 4, 7, -4},
				nums: []int32{0, 0, 0, 4, 4, -1, -1, -4, -7},
			},
			want: false,
		},
		{
			name: "TestInInt32s s nil",
			args: args{
				l:    []int32{0, 0, -1, 4, 7, -4},
				nums: nil,
			},
			want: true,
		},
		{
			name: "TestInInt32s both nil",
			args: args{
				l:    nil,
				nums: nil,
			},
			want: true,
		},
		{
			name: "TestInInt32s nums nil",
			args: args{
				l:    nil,
				nums: []int32{0, 0, 0, 4, 4, -1, -1, -4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InInt32s(tt.args.l, tt.args.nums...); got != tt.want {
				t.Errorf("InInt32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInUint64s(t *testing.T) {
	type args struct {
		l    []uint64
		nums []uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestInUint64s 1",
			args: args{
				l:    []uint64{0, 0, 4, 7},
				nums: []uint64{0, 0, 0, 4},
			},
			want: true,
		},
		{
			name: "TestInUint64s 2",
			args: args{
				l:    []uint64{0, 0, 4, 7},
				nums: []uint64{0, 0, 0, 4, 2},
			},
			want: false,
		},
		{
			name: "TestInUint64s s nil",
			args: args{
				l:    []uint64{0, 0, 4, 7},
				nums: nil,
			},
			want: true,
		},
		{
			name: "TestInUint64s both nil",
			args: args{
				l:    nil,
				nums: nil,
			},
			want: true,
		},
		{
			name: "TestInUint64s nums nil",
			args: args{
				l:    nil,
				nums: []uint64{0, 0, 0, 4, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InUint64s(tt.args.l, tt.args.nums...); got != tt.want {
				t.Errorf("InUint64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInUint32s(t *testing.T) {
	type args struct {
		l    []uint32
		nums []uint32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestInUint32s 1",
			args: args{
				l:    []uint32{0, 0, 4, 7},
				nums: []uint32{0, 0, 0, 4},
			},
			want: true,
		},
		{
			name: "TestInUint32s 2",
			args: args{
				l:    []uint32{0, 0, 4, 7},
				nums: []uint32{0, 0, 0, 4, 2},
			},
			want: false,
		},
		{
			name: "TestInUint32s s nil",
			args: args{
				l:    []uint32{0, 0, 4, 7},
				nums: nil,
			},
			want: true,
		},
		{
			name: "TestInUint32s both nil",
			args: args{
				l:    nil,
				nums: nil,
			},
			want: true,
		},
		{
			name: "TestInUint32s nums nil",
			args: args{
				l:    nil,
				nums: []uint32{0, 0, 0, 4, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InUint32s(tt.args.l, tt.args.nums...); got != tt.want {
				t.Errorf("InUint32s() = %v, want %v", got, tt.want)
			}
		})
	}
}
