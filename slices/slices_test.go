package slices

import (
	"reflect"
	"testing"
)

func Test_dropEmpty(t *testing.T) {
	type args struct {
		v []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"one", args{v: []string{"", "a", "", "b", "", "", "c", ""}}, []string{"a", "b", "c"}},
		{"two", args{v: []string{}}, nil},
		{"three", args{v: []string{"a", "b", "c"}}, []string{"a", "b", "c"}},
		{"three", args{v: []string{"", "b", "c"}}, []string{"b", "c"}},
		{"three", args{v: []string{"a", "b", ""}}, []string{"a", "b"}},
		{"four", args{v: []string{"a"}}, []string{"a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StrDropEmpty(tt.args.v)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want: %+v, got: %+v", tt.want, got)
			}
		})
	}
}
