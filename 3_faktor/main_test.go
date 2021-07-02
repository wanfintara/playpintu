package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findFactor(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "t1: Factor of 6",
			args: args{
				n: 6,
			},
			want: []int{1, 2, 3, 6},
		},
		{
			name: "t2: Factor of 20",
			args: args{
				n: 20,
			},
			want: []int{1, 2, 4, 5, 10, 20},
		},
		{
			name: "t3: Factor of 16",
			args: args{
				n: 16,
			},
			want: []int{1, 2, 4, 8, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// The function is not sorted yet
			got := findFactor(tt.args.n)
			sort.Ints(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findAll6FactorN(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		// want map[int][]int
		wantLength int
	}{
		{
			name: "t1: H(128)",
			args: args{
				n: 128,
			},
			wantLength: 19,
		},
		{
			name: "t2: H(1024)",
			args: args{
				n: 1024,
			},
			wantLength: 112,
		},
		{
			name: "t1: H(16384)",
			args: args{
				n: 16384,
			},
			wantLength: 1168,
		},
		{
			name: "t1: H(262144)",
			args: args{
				n: 262144,
			},
			wantLength: 13208,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAll6FactorN(tt.args.n)
			lengthGot := len(got)

			if !reflect.DeepEqual(lengthGot, tt.wantLength) {
				t.Errorf("length findAll6FactorN() = %v, want %v", lengthGot, tt.wantLength)
			}
		})
	}
}
