package slice

import (
	"reflect"
	"testing"
)

func TestMergeInt32s(t *testing.T) {
	type args struct {
		l [][]int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "TestMergeInt32s",
			args: args{l: [][]int32{
				{7, -1, 6, 8, 6, -4, 0},
				nil,
				{4, 6, 8, 2},
			}},
			want: []int32{7, -1, 6, 8, 6, -4, 0, 4, 6, 8, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeInt32s(tt.args.l...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeInt32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeInt64s(t *testing.T) {
	type args struct {
		l [][]int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "TestMergeInt64s",
			args: args{l: [][]int64{
				{7, -1, 6, 8, 6, -4, 0},
				nil,
				{4, 6, 8, 2},
			}},
			want: []int64{7, -1, 6, 8, 6, -4, 0, 4, 6, 8, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeInt64s(tt.args.l...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeInt64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeStrings(t *testing.T) {
	type args struct {
		l [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestMergeStrings",
			args: args{l: [][]string{
				{"time", "people", "time", ""},
				{""},
				nil,
				{"new", "time"},
			}},
			want: []string{"time", "people", "time", "", "", "new", "time"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeStrings(tt.args.l...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeUint32s(t *testing.T) {
	type args struct {
		l [][]uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "TestMergeUint32s",
			args: args{l: [][]uint32{
				{7, 6, 8, 6, 0},
				nil,
				{0},
			}},
			want: []uint32{7, 6, 8, 6, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeUint32s(tt.args.l...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeUint32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeUint64s(t *testing.T) {
	type args struct {
		l [][]uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "TestMergeUint64s",
			args: args{l: [][]uint64{
				{7, 6, 8, 6, 0},
				nil,
				{0},
			}},
			want: []uint64{7, 6, 8, 6, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeUint64s(tt.args.l...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeUint64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueMergeInt32s(t *testing.T) {
	type args struct {
		l [][]int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "TestUniqueMergeInt32s",
			args: args{l: [][]int32{
				{7, -1, 6, 8, 6, -4, 0},
				nil,
				{4, 6, 8, 2},
			}},
			want: []int32{7, -1, 6, 8, -4, 0, 4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueMergeInt32s(tt.args.l...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueMergeInt32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueMergeInt64s(t *testing.T) {
	type args struct {
		l [][]int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "TestUniqueMergeInt64s",
			args: args{l: [][]int64{
				{7, -1, 6, 8, 6, -4, 0},
				nil,
				{4, 6, 8, 2},
			}},
			want: []int64{7, -1, 6, 8, -4, 0, 4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueMergeInt64s(tt.args.l...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueMergeInt64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueMergeStrings(t *testing.T) {
	type args struct {
		l [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestUniqueMergeStrings",
			args: args{l: [][]string{
				{"time", "people", "time", ""},
				{""},
				nil,
				{"new", "time"},
			}},
			want: []string{"time", "people", "", "new"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueMergeStrings(tt.args.l...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueMergeStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueMergeUint32s(t *testing.T) {
	type args struct {
		l [][]uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "TestUniqueMergeUint32s",
			args: args{l: [][]uint32{
				{7, 6, 8, 6, 0},
				nil,
				{0},
			}},
			want: []uint32{7, 6, 8, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueMergeUint32s(tt.args.l...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueMergeUint32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueMergeUint64s(t *testing.T) {
	type args struct {
		l [][]uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "TestUniqueMergeUint64s",
			args: args{l: [][]uint64{
				{7, 6, 8, 6, 0},
				nil,
				{0},
			}},
			want: []uint64{7, 6, 8, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueMergeUint64s(tt.args.l...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueMergeUint64s() = %v, want %v", got, tt.want)
			}
		})
	}
}
