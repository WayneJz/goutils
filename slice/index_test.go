package slice

import (
	"reflect"
	"testing"
)

func TestIndexAllInt32s(t *testing.T) {
	type args struct {
		l    []int32
		nums []int32
	}
	tests := []struct {
		name string
		args args
		want map[int32]int
	}{
		{
			name: "TestIndexAllInt32s",
			args: args{
				l:    []int32{7, -9, 2, 4, 0, 2, -1},
				nums: []int32{-8, 2, 0, 2, 7},
			},
			want: map[int32]int{
				-8: -1,
				2:  2,
				0:  4,
				7:  0,
			},
		},
		{
			name: "TestIndexAllInt32s nil l",
			args: args{
				l:    nil,
				nums: []int32{-8, 2, 0, 2, 7},
			},
			want: map[int32]int{
				-8: -1,
				2:  -1,
				0:  -1,
				7:  -1,
			},
		},
		{
			name: "TestIndexAllInt32s nil nums",
			args: args{
				l:    []int32{7, -9, 2, 4, 0, 2, -1},
				nums: nil,
			},
			want: map[int32]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexAllInt32s(tt.args.l, tt.args.nums...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexAllInt32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexAllInt64s(t *testing.T) {
	type args struct {
		l    []int64
		nums []int64
	}
	tests := []struct {
		name string
		args args
		want map[int64]int
	}{
		{
			name: "TestIndexAllInt64s",
			args: args{
				l:    []int64{7, -9, 2, 4, 0, 2, -1},
				nums: []int64{-8, 2, 0, 2, 7},
			},
			want: map[int64]int{
				-8: -1,
				2:  2,
				0:  4,
				7:  0,
			},
		},
		{
			name: "TestIndexAllInt64s nil l",
			args: args{
				l:    nil,
				nums: []int64{-8, 2, 0, 2, 7},
			},
			want: map[int64]int{
				-8: -1,
				2:  -1,
				0:  -1,
				7:  -1,
			},
		},
		{
			name: "TestIndexAllInt64s nil nums",
			args: args{
				l:    []int64{7, -9, 2, 4, 0, 2, -1},
				nums: nil,
			},
			want: map[int64]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexAllInt64s(tt.args.l, tt.args.nums...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexAllInt64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexAllStrings(t *testing.T) {
	type args struct {
		l []string
		s []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "TestIndexAllStrings",
			args: args{
				l: []string{"empty", "", "hehe", "good", ""},
				s: []string{"", "hehe", ""},
			},
			want: map[string]int{
				"":     1,
				"hehe": 2,
			},
		},
		{
			name: "TestIndexAllStrings nil l",
			args: args{
				l: nil,
				s: []string{"", "hehe", ""},
			},
			want: map[string]int{
				"":     -1,
				"hehe": -1,
			},
		},
		{
			name: "TestIndexAllStrings nil nums",
			args: args{
				l: []string{"empty", "", "hehe", "good", ""},
				s: nil,
			},
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexAllStrings(tt.args.l, tt.args.s...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexAllStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexAllUint32s(t *testing.T) {
	type args struct {
		l    []uint32
		nums []uint32
	}
	tests := []struct {
		name string
		args args
		want map[uint32]int
	}{
		{
			name: "TestIndexAllUint32s",
			args: args{
				l:    []uint32{7, 0, 2, 4, 0},
				nums: []uint32{2, 0, 2, 7},
			},
			want: map[uint32]int{
				2: 2,
				0: 1,
				7: 0,
			},
		},
		{
			name: "TestIndexAllUint32s nil l",
			args: args{
				l:    nil,
				nums: []uint32{2, 0, 2, 7},
			},
			want: map[uint32]int{
				2: -1,
				0: -1,
				7: -1,
			},
		},
		{
			name: "TestIndexAllUint32s nil nums",
			args: args{
				l:    []uint32{7, 0, 2, 4, 0},
				nums: nil,
			},
			want: map[uint32]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexAllUint32s(tt.args.l, tt.args.nums...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexAllUint32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexAllUint64s(t *testing.T) {
	type args struct {
		l    []uint64
		nums []uint64
	}
	tests := []struct {
		name string
		args args
		want map[uint64]int
	}{
		{
			name: "TestIndexAllUint64s",
			args: args{
				l:    []uint64{7, 0, 2, 4, 0},
				nums: []uint64{2, 0, 2, 7},
			},
			want: map[uint64]int{
				2: 2,
				0: 1,
				7: 0,
			},
		},
		{
			name: "TestIndexAllUint64s nil l",
			args: args{
				l:    nil,
				nums: []uint64{2, 0, 2, 7},
			},
			want: map[uint64]int{
				2: -1,
				0: -1,
				7: -1,
			},
		},
		{
			name: "TestIndexAllUint64s nil nums",
			args: args{
				l:    []uint64{7, 0, 2, 4, 0},
				nums: nil,
			},
			want: map[uint64]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexAllUint64s(tt.args.l, tt.args.nums...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexAllUint64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexInt32s(t *testing.T) {
	type args struct {
		l   []int32
		num int32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestIndexInt32s",
			args: args{
				l:   []int32{7, -9, 2, 4, 0, 2, -1},
				num: 2,
			},
			want: 2,
		},
		{
			name: "TestIndexInt32s nil l",
			args: args{
				l:   nil,
				num: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexInt32s(tt.args.l, tt.args.num); got != tt.want {
				t.Errorf("IndexInt32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexInt64s(t *testing.T) {
	type args struct {
		l   []int64
		num int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestIndexInt64s",
			args: args{
				l:   []int64{7, -9, 2, 4, 0, 2, -1},
				num: 2,
			},
			want: 2,
		},
		{
			name: "TestIndexInt64s nil l",
			args: args{
				l:   nil,
				num: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexInt64s(tt.args.l, tt.args.num); got != tt.want {
				t.Errorf("IndexInt64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexStrings(t *testing.T) {
	type args struct {
		l []string
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestIndexStrings",
			args: args{
				l: []string{"empty", "", "hehe", "good", ""},
				s: "",
			},
			want: 1,
		},
		{
			name: "TestIndexAllStrings nil l",
			args: args{
				l: nil,
				s: "",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexStrings(tt.args.l, tt.args.s); got != tt.want {
				t.Errorf("IndexStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexUint32s(t *testing.T) {
	type args struct {
		l   []uint32
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestIndexUint32s",
			args: args{
				l:   []uint32{7, 2, 4, 0, 2},
				num: 2,
			},
			want: 1,
		},
		{
			name: "TestIndexUint32s nil l",
			args: args{
				l:   nil,
				num: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexUint32s(tt.args.l, tt.args.num); got != tt.want {
				t.Errorf("IndexUint32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexUint64s(t *testing.T) {
	type args struct {
		l   []uint64
		num uint64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestIndexUint32s",
			args: args{
				l:   []uint64{7, 2, 4, 0, 2},
				num: 2,
			},
			want: 1,
		},
		{
			name: "TestIndexUint32s nil l",
			args: args{
				l:   nil,
				num: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexUint64s(tt.args.l, tt.args.num); got != tt.want {
				t.Errorf("IndexUint64s() = %v, want %v", got, tt.want)
			}
		})
	}
}
