package main

import (
	"math"

)


type Player struct {
	X      int
	Y      int
	Hunter bool
	Symbol string
	Crumbs []coords
	Color string
}
type coords struct {
	X int
	Y int
}


func (this *Player)breadcrumb(maze [][]string, turn int) (alteredMaze [][]string) {
	emptyCoords := coords{}
	index := turn % len(this.Crumbs)
	if this.Crumbs[index] != emptyCoords {
		// decrumb
		coord := this.Crumbs[index]
		maze[coord.X][coord.Y] = emptyTile
	}
	// new crumb
	this.Crumbs[index] = coords{X:  this.X, Y: this.Y}
	maze[this.X][this.Y] = this.Color + emptyTile + TermReset
	return maze
}

// move the player around
func (this *Player) move(maze [][]string, dir int, turn int) (newMaze [][]string) {
	// println(this.X, this.Y)
	newX := this.X + dirs[dir][0]
	newY := this.Y + dirs[dir][1]
	// if maze[newX][newY] == emptyTile {
		mazeMutex.Lock()
		maze[this.X][this.Y] = emptyTile
		maze = this.breadcrumb(maze, turn)
		this.X = newX + dirs[dir][0]
		this.Y = newY + dirs[dir][1]
		maze[this.X][this.Y] = this.Symbol
		mazeMutex.Unlock()

	// }

	return maze
}
// places the player, also contains all the defaults
func (this *Player) placePlayer(maze [][]string) (newMaze [][]string) {
	for maze[this.X][this.Y] != emptyTile {
		if maze[this.X][this.Y+1] == emptyTile {
			this.Y += 1
		} else {
			this.X += 1
		}
	}
	if len(this.Crumbs) == 0 {
		this.Crumbs = make([]coords, 10)
	}
	if this.Color == "" {
		this.Color = TermGreen
	}
	mazeMutex.Lock()
	maze[this.X][this.Y] = this.Symbol
	mazeMutex.Unlock()
	return maze
}
// returns the viewport that this player sees
func (this *Player)viewPort(maze [][]string, size int) (alteredMaze [][]string) {
	yView := size + this.Y
	xView := size + this.X
	nyView := this.Y - size
	nxView := this.X - size
	// println(yView, nyView)
	if nyView < 0 {
		yView += int(math.Abs(float64(nyView)))
		nyView = 0
	} else if yView > len(maze[this.X]) {
		nyView += len(maze[this.X]) - int(math.Abs(float64(yView)))
		yView = len(maze[this.X]) 
	}
	if nxView < 0 {
		xView += int(math.Abs(float64(nxView)))
		nxView = 0
	} else if xView > len(maze) {
		nxView += len(maze) - int(math.Abs(float64(xView)))
		xView = len(maze)
	}
	for i := 0; nxView + i < xView; i++ {
		var tempRow []string
		for j := 0; nyView + j < yView; j++ {
			tempRow = append(tempRow, maze[nxView + i][nyView + j])
		}
		alteredMaze = append(alteredMaze, tempRow)
	}
	return alteredMaze
}
