package maps

import (
	"reflect"
	"testing"
)

func TestStringMapInOrder(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want [][]interface{}
	}{
		{
			name: "TestStringMapInOrder",
			args: args{
				m: map[string]interface{}{
					"numeric":   3,
					"anystring": "str",
					"slice":     []int64{1, 2},
					"anymap":    map[string]string{"m": "map"},
				},
			},
			want: [][]interface{}{
				{"anymap", map[string]string{"m": "map"}},
				{"anystring", "str"},
				{"numeric", 3},
				{"slice", []int64{1, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringMapInOrder(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringMapInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64MapInOrder(t *testing.T) {
	type args struct {
		m map[int64]interface{}
	}
	tests := []struct {
		name string
		args args
		want [][]interface{}
	}{
		{
			name: "TestInt64MapInOrder",
			args: args{
				m: map[int64]interface{}{
					7:  3,
					1:  "str",
					3:  []int64{1, 2},
					27: map[string]string{"m": "map"},
				},
			},
			want: [][]interface{}{
				{int64(1), "str"},
				{int64(3), []int64{1, 2}},
				{int64(7), 3},
				{int64(27), map[string]string{"m": "map"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64MapInOrder(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64MapInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32MapInOrder(t *testing.T) {
	type args struct {
		m map[int32]interface{}
	}
	tests := []struct {
		name string
		args args
		want [][]interface{}
	}{
		{
			name: "TestInt32MapInOrder",
			args: args{
				m: map[int32]interface{}{
					7:  3,
					1:  "str",
					3:  []int64{1, 2},
					27: map[string]string{"m": "map"},
				},
			},
			want: [][]interface{}{
				{int32(1), "str"},
				{int32(3), []int64{1, 2}},
				{int32(7), 3},
				{int32(27), map[string]string{"m": "map"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32MapInOrder(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32MapInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64MapInOrder(t *testing.T) {
	type args struct {
		m map[uint64]interface{}
	}
	tests := []struct {
		name string
		args args
		want [][]interface{}
	}{
		{
			name: "TestUint64MapInOrder",
			args: args{
				m: map[uint64]interface{}{
					7:  3,
					1:  "str",
					3:  []int64{1, 2},
					27: map[string]string{"m": "map"},
				},
			},
			want: [][]interface{}{
				{uint64(1), "str"},
				{uint64(3), []int64{1, 2}},
				{uint64(7), 3},
				{uint64(27), map[string]string{"m": "map"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64MapInOrder(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64MapInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32MapInOrder(t *testing.T) {
	type args struct {
		m map[uint32]interface{}
	}
	tests := []struct {
		name string
		args args
		want [][]interface{}
	}{
		{
			name: "TestUint32MapInOrder",
			args: args{
				m: map[uint32]interface{}{
					7:  3,
					1:  "str",
					3:  []int64{1, 2},
					27: map[string]string{"m": "map"},
				},
			},
			want: [][]interface{}{
				{uint32(1), "str"},
				{uint32(3), []int64{1, 2}},
				{uint32(7), 3},
				{uint32(27), map[string]string{"m": "map"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32MapInOrder(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32MapInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
