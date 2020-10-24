package maps

import (
	"fmt"
	"testing"
)

func TestStringToJSONMap(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "TestStringToJSONMap",
			args: args{
				str: `{"id":123,"addr":"testaddr","tag":[0,1],"embedded":{"test":"ok"}}`,
			},
			want: map[string]interface{}{
				"id":       123,
				"addr":     "testaddr",
				"tag":      []interface{}{0, 1},
				"embedded": map[string]interface{}{"test": "ok"},
			},
			wantErr: false,
		},
		{
			name: "TestStringToJSONMap unmarshal failed",
			args: args{
				str: `haha`,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToJSONMap(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToJSONMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPrint, wantPrint := fmt.Sprintf("%#v", got), fmt.Sprintf("%#v", tt.want); gotPrint != wantPrint {
				t.Errorf("StringToJSONMap() gotPrint = %v, wantPrint = %v", gotPrint, wantPrint)
				return
			}
		})
	}
}

func TestMustStringToJSONMap(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "TestMustStringToJSONMap",
			args: args{
				str: `{"id":123,"addr":"testaddr","tag":[0,1],"embedded":{"test":"ok"}}`,
			},
			want: map[string]interface{}{
				"id":       123,
				"addr":     "testaddr",
				"tag":      []interface{}{0, 1},
				"embedded": map[string]interface{}{"test": "ok"},
			},
		},
		{
			name: "TestMustStringToJSONMap unmarshal failed",
			args: args{
				str: `haha`,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MustStringToJSONMap(tt.args.str)
			if gotPrint, wantPrint := fmt.Sprintf("%#v", got), fmt.Sprintf("%#v", tt.want); gotPrint != wantPrint {
				t.Errorf("MustStringToJSONMap() gotPrint = %v, wantPrint = %v", gotPrint, wantPrint)
				return
			}
		})
	}
}

func TestStringsToJSONMaps(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []map[string]interface{}
		wantErr bool
	}{
		{
			name: "TestStringsToJSONMaps",
			args: args{
				strs: []string{`{"id":123,"addr":"testaddr","tag":[0,1],"embedded":{"test":"ok"}}`},
			},
			wantRes: []map[string]interface{}{
				{
					"id":       123,
					"addr":     "testaddr",
					"tag":      []interface{}{0, 1},
					"embedded": map[string]interface{}{"test": "ok"},
				},
			},
			wantErr: false,
		},
		{
			name: "TestStringsToJSONMaps unmarshal failed",
			args: args{
				strs: []string{`haha`},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringsToJSONMaps(tt.args.strs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringsToJSONMaps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPrint, wantPrint := fmt.Sprintf("%#v", got), fmt.Sprintf("%#v", tt.wantRes); gotPrint != wantPrint {
				t.Errorf("StringsToJSONMaps() gotPrint = %v, wantPrint = %v", gotPrint, wantPrint)
				return
			}
		})
	}
}

func TestMustStringsToJSONMaps(t *testing.T) {
	type args struct {
		strs []string
	}
	var m map[string]interface{}
	tests := []struct {
		name    string
		args    args
		wantRes []map[string]interface{}
	}{
		{
			name: "TestMustStringsToJSONMaps",
			args: args{
				strs: []string{`{"id":123,"addr":"testaddr","tag":[0,1],"embedded":{"test":"ok"}}`},
			},
			wantRes: []map[string]interface{}{
				{
					"id":       123,
					"addr":     "testaddr",
					"tag":      []interface{}{0, 1},
					"embedded": map[string]interface{}{"test": "ok"},
				},
			},
		},
		{
			name: "TestMustStringsToJSONMaps unmarshal failed",
			args: args{
				strs: []string{`haha`},
			},
			wantRes: []map[string]interface{}{m},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MustStringsToJSONMaps(tt.args.strs...)
			if gotPrint, wantPrint := fmt.Sprintf("%#v", got), fmt.Sprintf("%#v", tt.wantRes); gotPrint != wantPrint {
				t.Errorf("MustStringsToJSONMaps() gotPrint = %v, wantPrint = %v", gotPrint, wantPrint)
				return
			}
		})
	}
}
