package slice

import (
	"testing"
)

func TestReverseStrings(t *testing.T) {
	type args struct {
		l []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseStrings(tt.args.l)
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseInt64s(tt.args.l)
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseInt32s(tt.args.l)
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseUint64s(tt.args.l)
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseUint32s(tt.args.l)
		})
	}
}
