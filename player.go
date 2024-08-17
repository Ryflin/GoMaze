package main

import (
	"math"
	"time"
)

// TODO remove the mutexes as this is no longer async
// when doing the async use a temporary variable to store the new info and swap in when mutex is unlocked.

type Player struct {
	X      int
	Y      int
	Hunter bool
	Symbol string
	Crumbs []coords
	Color  string
}
type coords struct {
	X int
	Y int
}

func (player *Player) breadcrumb(maze [][]string, turn int) (alteredMaze [][]string) {
	emptyCoords := coords{}
	index := turn % len(player.Crumbs)
	currentCoords := coords{X: player.X, Y: player.Y}
	// need to also remove duplicates
	for i, val := range player.Crumbs {
		if val == currentCoords {
			player.Crumbs[i] = emptyCoords
		}
	}
	if player.Crumbs[index] != emptyCoords {
		// decrumb
		coord := player.Crumbs[index]
		maze[coord.X][coord.Y] = emptyTile
	}

	// new crumb
	player.Crumbs[index] = currentCoords
	maze[player.X][player.Y] = player.Color + emptyTile + TermReset
	return maze
}

// move the player around
func (player *Player) move(maze [][]string, dir int, turn int) (newMaze [][]string) {
	// println(this.X, this.Y)
	newX := player.X + dirs[dir][0]
	newY := player.Y + dirs[dir][1]
	if maze[newX][newY] == emptyTile {
	// if true {
		maze[player.X][player.Y] = emptyTile
		mazeMutex.Lock()
		maze = player.breadcrumb(maze, turn)
		player.X = newX
		player.Y = newY
		maze[player.X][player.Y] = player.Symbol
		render(player.viewPort(maze, viewPortSize))
		mazeMutex.Unlock()
		time.Sleep(time.Second / 15)
		mazeMutex.Lock()
		maze[newX][newY] = emptyTile
		player.X = newX + dirs[dir][0]
		player.Y = newY + dirs[dir][1]
		maze[player.X][player.Y] = player.Symbol
		mazeMutex.Unlock()
	}

	return maze
}

// places the player, also contains all the defaults
func (player *Player) placePlayer(maze [][]string) (newMaze [][]string) {
	for maze[player.X][player.Y] != emptyTile {
		if maze[player.X][player.Y+1] == emptyTile {
			player.Y += 1
		} else {
			player.X += 1
		}
	}
	if len(player.Crumbs) == 0 {
		player.Crumbs = make([]coords, 20)
	}
	if player.Color == "" {
		player.Color = CrumbBlue
	}
	mazeMutex.Lock()
	maze[player.X][player.Y] = player.Symbol
	mazeMutex.Unlock()
	return maze
}

// returns the viewport that this player sees
func (player *Player) viewPort(maze [][]string, size int) (alteredMaze [][]string) {
	yView := size + player.Y
	xView := size + player.X
	nyView := player.Y - size
	nxView := player.X - size
	// println(yView, nyView)
	if nyView < 0 {
		yView += int(math.Abs(float64(nyView)))
		nyView = 0
	} else if yView > len(maze[player.X]) {
		nyView += len(maze[player.X]) - int(math.Abs(float64(yView)))
		yView = len(maze[player.X])
	}
	if nxView < 0 {
		xView += int(math.Abs(float64(nxView)))
		nxView = 0
	} else if xView > len(maze) {
		nxView += len(maze) - int(math.Abs(float64(xView)))
		xView = len(maze)
	}
	for i := 0; nxView+i < xView; i++ {
		var tempRow []string
		for j := 0; nyView+j < yView; j++ {
			tempRow = append(tempRow, maze[nxView+i][nyView+j])
		}
		alteredMaze = append(alteredMaze, tempRow)
	}
	return alteredMaze
}
