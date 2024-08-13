package main

import (
	"errors"
	"fmt"
)

var (
	dirs = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
)

const (
	MaxSize = 38
	MinSize = 3
	wallRep = "#"
	sqrRep  = "•"
	escUp   = "\033[F"
)

func newError(errorString string) error {
	return errors.New(errorString)
}

var (
	ErrMazeTooBig   = fmt.Errorf("Error provided maze size too big, current max is %d", MaxSize)
	ErrMazeTooSmall = fmt.Errorf("Error provided maze size is too small. current min is %d", MinSize)
)
