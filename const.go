package main

import (
	"errors"
	"fmt"
	"sync"
)

var (
 // this is just a comment that I want to put in because I want to. 
	dirs = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	// This is for putting required values you don't want to use into a bucket. bad practice I know but makes other things more convienient
	inputs    = map[string]int{"w": 0, "a": 1, "r": 2, "s": 3}
	mazeMutex sync.Mutex
	wallEnd   = HighlightWhite + "   " + TermReset
	biomeList = []string{HighlightBlue, HighlightCyan, HighlightGreen, HighlightMagenta, HighlightRed, HighlightYellow}
	winState = 0
	enemy = Hunter{}
)

const (
	exit             = "EXT"
	biomeWidth       = 10
	viewPortSize     = 10
	TermReset        = "\033[0m"
	TermRed          = "\033[31m"
	TermGreen        = "\033[32m"
	TermYellow       = "\033[33m"
	TermBlue         = "\033[34m"
	TermPurple       = "\033[35m"
	TermCyan         = "\033[36m"
	TermGray         = "\033[37m"
	TermWhite        = "\033[97m"
	CrumbBlue        = "\033[38;2;0;0;255;48;2;10;10;125m"
	HighlightWhite   = "\033[37;47m"
	HighlightRed     = "\033[31;41m"
	HighlightGreen   = "\033[32;42m"
	HighlightYellow  = "\033[33;43m"
	HighlightBlue    = "\033[34;44m"
	HighlightMagenta = "\033[35;45m"
	HighlightCyan    = "\033[36;46m"
	MaxSize          = 38
	MinSize          = 3
	emptyTile        = " • "
	escUp            = "\033[F"
	moveStep         = int(100)
	biomeSize        = 10
)

func newError(errorString string) error {
	return errors.New(errorString)
}

var (
	ErrMazeTooBig   = fmt.Errorf("Error provided maze size too big, current max is %d", MaxSize)
	ErrMazeTooSmall = fmt.Errorf("Error provided maze size is too small. current min is %d", MinSize)
)
