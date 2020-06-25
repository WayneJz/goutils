package slice

import (
	"reflect"
	"testing"
)

func TestApplyInt32s(t *testing.T) {
	type args struct {
		l         []int32
		applyFunc func(n int32) int32
	}
	tests := []struct {
		name      string
		args      args
		wantSlice []int32
	}{
		{
			name: "TestApplyInt32s multiply by 2",
			args: args{
				l: []int32{9, 5, 3, 6, 7},
				applyFunc: func(n int32) int32 {
					return n * 2
				},
			},
			wantSlice: []int32{18, 10, 6, 12, 14},
		},
		{
			name: "TestApplyInt32s divided by 2 (NOTE: AUTO CASTED)",
			args: args{
				l: []int32{9, 5, 3, 6, 7},
				applyFunc: func(n int32) int32 {
					return n / 2
				},
			},
			wantSlice: []int32{4, 2, 1, 3, 3},
		},
		{
			name: "TestApplyInt32s in nil slice",
			args: args{
				l: nil,
				applyFunc: func(n int32) int32 {
					return n / 2
				},
			},
			wantSlice: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ApplyInt32s(tt.args.l, tt.args.applyFunc)
			if !reflect.DeepEqual(tt.args.l, tt.wantSlice) {
				t.Errorf("after apply = %+v, want %+v", tt.args.l, tt.wantSlice)
			}
		})
	}
}

func TestApplyInt64s(t *testing.T) {
	type args struct {
		l         []int64
		applyFunc func(n int64) int64
	}
	tests := []struct {
		name      string
		args      args
		wantSlice []int64
	}{
		{
			name: "TestApplyInt64s multiply by 2",
			args: args{
				l: []int64{9, 5, 3, 6, 7},
				applyFunc: func(n int64) int64 {
					return n * 2
				},
			},
			wantSlice: []int64{18, 10, 6, 12, 14},
		},
		{
			name: "TestApplyInt64s divided by 2 (NOTE: AUTO CASTED)",
			args: args{
				l: []int64{9, 5, 3, 6, 7},
				applyFunc: func(n int64) int64 {
					return n / 2
				},
			},
			wantSlice: []int64{4, 2, 1, 3, 3},
		},
		{
			name: "TestApplyInt64s in nil slice",
			args: args{
				l: nil,
				applyFunc: func(n int64) int64 {
					return n / 2
				},
			},
			wantSlice: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ApplyInt64s(tt.args.l, tt.args.applyFunc)
			if !reflect.DeepEqual(tt.args.l, tt.wantSlice) {
				t.Errorf("after apply = %+v, want %+v", tt.args.l, tt.wantSlice)
			}
		})
	}
}

func TestApplyStrings(t *testing.T) {
	type args struct {
		l         []string
		applyFunc func(s string) string
	}
	tests := []struct {
		name      string
		args      args
		wantSlice []string
	}{
		{
			name: "TestApplyStrings extend each string hello world at front",
			args: args{
				l: []string{"nice", "cool", "luck"},
				applyFunc: func(s string) string {
					return "hello world " + s
				},
			},
			wantSlice: []string{"hello world nice", "hello world cool", "hello world luck"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ApplyStrings(tt.args.l, tt.args.applyFunc)
			if !reflect.DeepEqual(tt.args.l, tt.wantSlice) {
				t.Errorf("after apply = %+v, want %+v", tt.args.l, tt.wantSlice)
			}
		})
	}
}

func TestApplyUint32s(t *testing.T) {
	type args struct {
		l         []uint32
		applyFunc func(n uint32) uint32
	}
	tests := []struct {
		name      string
		args      args
		wantSlice []uint32
	}{
		{
			name: "TestApplyUint32s multiply by 2",
			args: args{
				l: []uint32{9, 5, 3, 6, 7},
				applyFunc: func(n uint32) uint32 {
					return n * 2
				},
			},
			wantSlice: []uint32{18, 10, 6, 12, 14},
		},
		{
			name: "TestApplyUint32s divided by 2 (NOTE: AUTO CASTED)",
			args: args{
				l: []uint32{9, 5, 3, 6, 7},
				applyFunc: func(n uint32) uint32 {
					return n / 2
				},
			},
			wantSlice: []uint32{4, 2, 1, 3, 3},
		},
		{
			name: "TestApplyUint32s in nil slice",
			args: args{
				l: nil,
				applyFunc: func(n uint32) uint32 {
					return n / 2
				},
			},
			wantSlice: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ApplyUint32s(tt.args.l, tt.args.applyFunc)
			if !reflect.DeepEqual(tt.args.l, tt.wantSlice) {
				t.Errorf("after apply = %+v, want %+v", tt.args.l, tt.wantSlice)
			}
		})
	}
}

func TestApplyUint64s(t *testing.T) {
	type args struct {
		l         []uint64
		applyFunc func(n uint64) uint64
	}
	tests := []struct {
		name      string
		args      args
		wantSlice []uint64
	}{
		{
			name: "TestApplyUint64s multiply by 2",
			args: args{
				l: []uint64{9, 5, 3, 6, 7},
				applyFunc: func(n uint64) uint64 {
					return n * 2
				},
			},
			wantSlice: []uint64{18, 10, 6, 12, 14},
		},
		{
			name: "TestApplyUint64s divided by 2 (NOTE: AUTO CASTED)",
			args: args{
				l: []uint64{9, 5, 3, 6, 7},
				applyFunc: func(n uint64) uint64 {
					return n / 2
				},
			},
			wantSlice: []uint64{4, 2, 1, 3, 3},
		},
		{
			name: "TestApplyUint64s in nil slice",
			args: args{
				l: nil,
				applyFunc: func(n uint64) uint64 {
					return n / 2
				},
			},
			wantSlice: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ApplyUint64s(tt.args.l, tt.args.applyFunc)
			if !reflect.DeepEqual(tt.args.l, tt.wantSlice) {
				t.Errorf("after apply = %+v, want %+v", tt.args.l, tt.wantSlice)
			}
		})
	}
}

func TestIApplyInt32s(t *testing.T) {
	type args struct {
		l         []int32
		applyFunc func(i int, n int32) int32
	}
	tests := []struct {
		name      string
		args      args
		wantSlice []int32
	}{
		{
			name: "TestIApplyInt32s multiply by the index",
			args: args{
				l: []int32{9, 5, 3, 6, 7},
				applyFunc: func(i int, n int32) int32 {
					return n * int32(i)
				},
			},
			wantSlice: []int32{0, 5, 6, 18, 28},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IApplyInt32s(tt.args.l, tt.args.applyFunc)
			if !reflect.DeepEqual(tt.args.l, tt.wantSlice) {
				t.Errorf("after apply = %+v, want %+v", tt.args.l, tt.wantSlice)
			}
		})
	}
}

func TestIApplyInt64s(t *testing.T) {
	type args struct {
		l         []int64
		applyFunc func(i int, n int64) int64
	}
	tests := []struct {
		name      string
		args      args
		wantSlice []int64
	}{
		{
			name: "TestIApplyInt64s multiply by the index",
			args: args{
				l: []int64{9, 5, 3, 6, 7},
				applyFunc: func(i int, n int64) int64 {
					return n * int64(i)
				},
			},
			wantSlice: []int64{0, 5, 6, 18, 28},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IApplyInt64s(tt.args.l, tt.args.applyFunc)
			if !reflect.DeepEqual(tt.args.l, tt.wantSlice) {
				t.Errorf("after apply = %+v, want %+v", tt.args.l, tt.wantSlice)
			}
		})
	}
}

func TestIApplyStrings(t *testing.T) {
	type args struct {
		l         []string
		applyFunc func(i int, s string) string
	}
	tests := []struct {
		name      string
		args      args
		wantSlice []string
	}{
		{
			name: "TestIApplyStrings extend each string hello world at front if index is odd number",
			args: args{
				l: []string{"nice", "cool", "luck"},
				applyFunc: func(i int, s string) string {
					if i%2 == 0 {
						return "hello world " + s
					}
					return s
				},
			},
			wantSlice: []string{"hello world nice", "cool", "hello world luck"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IApplyStrings(tt.args.l, tt.args.applyFunc)
			if !reflect.DeepEqual(tt.args.l, tt.wantSlice) {
				t.Errorf("after apply = %+v, want %+v", tt.args.l, tt.wantSlice)
			}
		})
	}
}

func TestIApplyUint32s(t *testing.T) {
	type args struct {
		l         []uint32
		applyFunc func(i int, n uint32) uint32
	}
	tests := []struct {
		name      string
		args      args
		wantSlice []uint32
	}{
		{
			name: "TestIApplyUint32s multiply by the index",
			args: args{
				l: []uint32{9, 5, 3, 6, 7},
				applyFunc: func(i int, n uint32) uint32 {
					return n * uint32(i)
				},
			},
			wantSlice: []uint32{0, 5, 6, 18, 28},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IApplyUint32s(tt.args.l, tt.args.applyFunc)
			if !reflect.DeepEqual(tt.args.l, tt.wantSlice) {
				t.Errorf("after apply = %+v, want %+v", tt.args.l, tt.wantSlice)
			}
		})
	}
}

func TestIApplyUint64s(t *testing.T) {
	type args struct {
		l         []uint64
		applyFunc func(i int, n uint64) uint64
	}
	tests := []struct {
		name      string
		args      args
		wantSlice []uint64
	}{
		{
			name: "TestIApplyUint64s multiply by the index",
			args: args{
				l: []uint64{9, 5, 3, 6, 7},
				applyFunc: func(i int, n uint64) uint64 {
					return n * uint64(i)
				},
			},
			wantSlice: []uint64{0, 5, 6, 18, 28},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IApplyUint64s(tt.args.l, tt.args.applyFunc)
			if !reflect.DeepEqual(tt.args.l, tt.wantSlice) {
				t.Errorf("after apply = %+v, want %+v", tt.args.l, tt.wantSlice)
			}
		})
	}
}
