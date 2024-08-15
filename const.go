package main

import (
	"errors"
	"fmt"
	"sync"
)

var (
	dirs = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	// This is for putting required values you don't want to use into a bucket. bad practice I know but makes other things more convienient
	nullable any
	inputs   = map[string]int{"w": 0, "a": 1, "r": 2, "s": 3}
	mazeMutex sync.Mutex
	mazeGlobal [][]string
)

const (
	MaxSize   = 38
	MinSize   = 3
	wallRep   = "#"
	emptyTile = "•"
	escUp     = "\033[F"
	moveStep = int(100)
)

func newError(errorString string) error {
	return errors.New(errorString)
}

var (
	ErrMazeTooBig   = fmt.Errorf("Error provided maze size too big, current max is %d", MaxSize)
	ErrMazeTooSmall = fmt.Errorf("Error provided maze size is too small. current min is %d", MinSize)
)
