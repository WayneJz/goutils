package errors

import (
	"fmt"
	"testing"
)

func TestErrors_Add(t *testing.T) {
	type fields struct {
		msg []string
	}
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantStr string
	}{
		{
			name: "TestErrors_Add empty",
			fields: fields{
				msg: nil,
			},
			args:    args{err: fmt.Errorf("first error")},
			wantStr: "errors traceback: sequence=0, err=first error",
		},
		{
			name: "TestErrors_Add non-empty",
			fields: fields{
				msg: []string{"first error"},
			},
			args:    args{err: fmt.Errorf("second error")},
			wantStr: "errors traceback: sequence=0, err=first error\nerrors traceback: sequence=1, err=second error",
		},
		{
			name: "TestErrors_Add add nil error from empty",
			fields: fields{
				msg: nil,
			},
			args:    args{err: nil},
			wantStr: "",
		},
		{
			name: "TestErrors_Add add nil error from non-empty",
			fields: fields{
				msg: []string{"first error"},
			},
			args:    args{err: nil},
			wantStr: "errors traceback: sequence=0, err=first error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Errors{
				msg: tt.fields.msg,
			}
			e.Add(tt.args.err)
			if got := e.Error(); got != tt.wantStr {
				t.Errorf("Error() = %v, want %v", got, tt.wantStr)
			}
		})
	}
}

func TestErrors_Error(t *testing.T) {
	type fields struct {
		msg []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestErrors_Error single",
			fields: fields{
				msg: []string{"first error"},
			},
			want: "errors traceback: sequence=0, err=first error",
		},
		{
			name: "TestErrors_Error mutiple",
			fields: fields{
				msg: []string{"first error", "second error"},
			},
			want: "errors traceback: sequence=0, err=first error\nerrors traceback: sequence=1, err=second error",
		},
		{
			name: "TestErrors_Error none",
			fields: fields{
				msg: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Errors{
				msg: tt.fields.msg,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrors_Return(t *testing.T) {
	type fields struct {
		msg []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "TestErrors_Return single",
			fields: fields{
				msg: []string{"first error"},
			},
			wantErr: true,
		},
		{
			name: "TestErrors_Error mutiple",
			fields: fields{
				msg: []string{"first error", "second error"},
			},
			wantErr: true,
		},
		{
			name: "TestErrors_Error none",
			fields: fields{
				msg: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Errors{
				msg: tt.fields.msg,
			}
			if err := e.Return(); (err != nil) != tt.wantErr {
				t.Errorf("Return() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
