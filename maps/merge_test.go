package maps

import (
	"reflect"
	"testing"
)

func TestMergeStringMaps(t *testing.T) {
	type args struct {
		maps []map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "TestMergeStringMaps",
			args: args{
				maps: []map[string]interface{}{
					{"23": 123},
					{"1": "test"},
					{"7": []int32{234, 567}},
				},
			},
			want: map[string]interface{}{
				"23": 123,
				"1":  "test",
				"7":  []int32{234, 567},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeStringMaps(tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeStringMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeInt64Maps(t *testing.T) {
	type args struct {
		maps []map[int64]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[int64]interface{}
	}{
		{
			name: "TestMergeInt64Maps",
			args: args{
				maps: []map[int64]interface{}{
					{23: 123},
					{1: "test"},
					{7: []int32{234, 567}},
				},
			},
			want: map[int64]interface{}{
				23: 123,
				1:  "test",
				7:  []int32{234, 567},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeInt64Maps(tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeInt64Maps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeInt32Maps(t *testing.T) {
	type args struct {
		maps []map[int32]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[int32]interface{}
	}{
		{
			name: "TestMergeInt32Maps",
			args: args{
				maps: []map[int32]interface{}{
					{23: 123},
					{1: "test"},
					{7: []int32{234, 567}},
				},
			},
			want: map[int32]interface{}{
				23: 123,
				1:  "test",
				7:  []int32{234, 567},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeInt32Maps(tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeInt32Maps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeUint64Maps(t *testing.T) {
	type args struct {
		maps []map[uint64]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[uint64]interface{}
	}{
		{
			name: "TestMergeUint64Maps",
			args: args{
				maps: []map[uint64]interface{}{
					{23: 123},
					{1: "test"},
					{7: []int32{234, 567}},
				},
			},
			want: map[uint64]interface{}{
				23: 123,
				1:  "test",
				7:  []int32{234, 567},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeUint64Maps(tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeUint64Maps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeUint32Maps(t *testing.T) {
	type args struct {
		maps []map[uint32]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[uint32]interface{}
	}{
		{
			name: "TestMergeUint32Maps",
			args: args{
				maps: []map[uint32]interface{}{
					{23: 123},
					{1: "test"},
					{7: []int32{234, 567}},
				},
			},
			want: map[uint32]interface{}{
				23: 123,
				1:  "test",
				7:  []int32{234, 567},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeUint32Maps(tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeUint32Maps() = %v, want %v", got, tt.want)
			}
		})
	}
}
