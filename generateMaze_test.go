package main

import "testing"

func Test_drawMaze(t *testing.T) {
	type args struct {
		list []Edge
		size int
	}
	tests := []struct {
		name string
		args args
	}{
		{"blank test 1", args{[]Edge{{X: 0, Y: 0, Dir: 2}}, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			drawMaze(tt.args.list, tt.args.size)
		})
	}
}
