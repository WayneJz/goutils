package jsonutil

import (
	"reflect"
	"testing"
	"time"
)

func TestJSONTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		tr      JSONTime
		want    []byte
		wantErr bool
	}{
		{
			name:    "marshal test",
			tr:      JSONTime(time.Date(2022, time.July, 10, 10, 30, 10, 0, time.Local)),
			want:    []byte("\"2022-07-10 10:30:10\""),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tt.tr
			got, err := tr.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONTime.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONTime.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}
