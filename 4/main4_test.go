package main

import (
	"reflect"
	"testing"
)

func Test_checkSlice(t *testing.T) {
	type args struct {
		slice1 []string
		slice2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "без элементов",
			args: args{
				slice1: []string{},
				slice2: []string{"a", "b", "c"},
			},
			want: []string{},
		},
		{
			name: "с элементами",
			args: args{
				slice1: []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
				slice2: []string{"banana", "date", "fig"},
			},
			want: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name: "одинаковые элементы",
			args: args{
				slice1: []string{"a", "b", "c"},
				slice2: []string{"a", "b", "c"},
			},
			want: []string{},
		},
		{
			name: "без элементов во 2 слайсе",
			args: args{
				slice1: []string{"a", "b", "c"},
				slice2: []string{},
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSlice(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
