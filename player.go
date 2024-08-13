package main

type Player struct {
	X int
	Y int
	Hunter bool
	Symbol string
}

func (*Player)NewPlayer()