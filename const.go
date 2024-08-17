package main

import (
	"errors"
	"fmt"
	"sync"
)

var (
	dirs = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	// This is for putting required values you don't want to use into a bucket. bad practice I know but makes other things more convienient
	nullable   any
	inputs     = map[string]int{"w": 0, "a": 1, "r": 2, "s": 3}
	mazeMutex  sync.Mutex
	mazeGlobal [][]string
	wallEnd    = HighlightWhite + "   " + TermReset
)

const (
	biomeWidth     = 10
	viewPortSize   = 10
	TermReset      = "\033[0m"
	TermRed        = "\033[31m"
	TermGreen      = "\033[32m"
	TermYellow     = "\033[33m"
	TermBlue       = "\033[34m"
	TermPurple     = "\033[35m"
	TermCyan       = "\033[36m"
	TermGray       = "\033[37m"
	TermWhite      = "\033[97m"
	CrumbBlue      = "\033[38;2;0;0;255;48;2;10;10;125m"
	HighlightWhite = "\033[37;47m"
	HighlightRed   = "\033[31;41m"
	MaxSize        = 38
	MinSize        = 3
	emptyTile      = " • "
	escUp          = "\033[F"
	moveStep       = int(100)
)

func newError(errorString string) error {
	return errors.New(errorString)
}

var (
	ErrMazeTooBig   = fmt.Errorf("Error provided maze size too big, current max is %d", MaxSize)
	ErrMazeTooSmall = fmt.Errorf("Error provided maze size is too small. current min is %d", MinSize)
)
