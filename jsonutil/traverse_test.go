package jsonutil

import (
	"reflect"
	"testing"

	"github.com/tidwall/gjson"
)

// TestTraverse traverse test
func TestTraverse(t *testing.T) {

	testcase := `{
		"name": {"first": "Tom", "last": "Anderson"},
		"age":37,
		"children": ["Sara","Alex","Jack"],
		"fav.movie": "Deer Hunter",
		"friends": [
		  {"first": "James", "last": "Murphy"},
		  {"first": "Roger", "last": "Craig"}
		],
		"users":{
			"2313":{"name":"Sara"},
			"7839":{"name":"Andy"}
		}
	}`
	res := make(map[string]interface{})

	type args struct {
		b        string
		callback TraverseCallback
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Traverse",
			args: args{
				b: testcase,
				callback: func(key *Key, value *gjson.Result, path []byte) {
					res[string(path)] = value.Value()
				},
			},
			want: map[string]interface{}{
				"age":              float64(37),
				"children.0":       "Sara",
				"children.1":       "Alex",
				"children.2":       "Jack",
				"fav\\.movie":      "Deer Hunter",
				"friends.0.first":  "James",
				"friends.0.last":   "Murphy",
				"friends.1.first":  "Roger",
				"friends.1.last":   "Craig",
				"name.first":       "Tom",
				"name.last":        "Anderson",
				"users.:2313.name": "Sara",
				"users.:7839.name": "Andy",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Traverse(tt.args.b, tt.args.callback)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("Traverse() = %v, want %v", res, tt.want)
			}
		})
	}
}
