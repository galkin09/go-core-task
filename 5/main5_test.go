package main

import (
	"reflect"
	"testing"
)

func Test_checkSlices(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []int
	}{
		{
			name: "test1",
			args: args{
				a: []int{1, 2, 3},
				b: []int{4, 5, 6},
			},
			want:  false,
			want1: []int{},
		},
		{
			name: "test2",
			args: args{
				a: []int{65, 3, 58, 678, 64},
				b: []int{64, 2, 3, 43},
			},
			want:  true,
			want1: []int{64, 3},
		},
		{
			name: "test3",
			args: args{
				a: []int{1, 2, 3},
				b: []int{1, 2, 3},
			},
			want:  true,
			want1: []int{1, 2, 3},
		},
		{
			name: "test4",
			args: args{
				a: []int{1, 2, 3},
				b: []int{},
			},
			want:  false,
			want1: []int{},
		},
		{
			name: "test5",
			args: args{
				a: []int{},
				b: []int{},
			},
			want:  false,
			want1: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := checkSlices(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("checkSlices() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("checkSlices() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
