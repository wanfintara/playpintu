package main

import "testing"

func Test_findDiff(t *testing.T) {
	type args struct {
		ary []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "T1: Example 1",
			args: args{
				ary: []int{5, 4, 3, 2, 1},
			},
			want: 0,
		},
		{
			name: "T2: Example 2",
			args: args{
				ary: []int{3, 2, 1, 5, 6, 2},
			},
			want: 5,
		},
		{
			name: "T3: Custom",
			args: args{
				ary: []int{3, 2, 99, 5, 6, 2, 8, 14, 7, 100, 1},
			},
			want: 98,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiff(tt.args.ary); got != tt.want {
				t.Errorf("findDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
