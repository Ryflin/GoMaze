package main

import (
	"errors"
	"fmt"
)
const (
	MazeSizeLimit = 38
	MazeSizeMin = 3
)
func newError(errorString string) error {
	return errors.New(errorString)
}

var (
	ErrMazeTooBig = fmt.Errorf("Error provided maze size too big, current max is %d", MazeSizeLimit)
	ErrMazeTooSmall = fmt.Errorf("Error provided maze size is too small. current min is %d", MazeSizeMin)
)