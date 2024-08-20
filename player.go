package main

import (
	"math"
	"time"
)

// TODO remove the mutexes as this is no longer async
// when doing the async use a temporary variable to store the new info and swap in when mutex is unlocked.\

// player options so that the player can be initialized with defaults
type PlayerOptions func(*Player)

func withCoords(coordinates Coords) PlayerOptions {
	return func(p *Player) {
		p.X = coordinates.X
		p.Y = coordinates.Y
	}
}
func withSymbol(Symbol string) PlayerOptions {
	return func(p *Player) {
		p.Symbol = Symbol
	}
}
func withColor(Color string) PlayerOptions {
	return func(p *Player) {
		p.Color = Color
	}
}
func withCrumbs(Crumbs []Coords) PlayerOptions {
	return func(p *Player) {
		p.Crumbs = Crumbs
	}
}
func newPlayer(opts ...PlayerOptions) *Player {
	player := &Player{
		X:      1,
		Y:      2,
		Symbol: " R ",
		Crumbs: make([]Coords, 20),
		Color:  TermBlue,
	}
	for _, opt := range opts {
		opt(player)
	}
	return player
}

type Player struct {
	X      int `json:"X"`
	Y      int `json:"Y"`
	Symbol string 
	Crumbs []Coords `json:"Crumbs"`
	Color  string
}
type Coords struct {
	X int
	Y int
}

func (player *Player) breadcrumb(maze [][]string, turn int) (alteredMaze [][]string) {
	emptyCoords := Coords{}
	index := turn % len(player.Crumbs)
	currentCoords := Coords{X: player.X, Y: player.Y}
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
	if maze[player.X][player.Y] != player.Symbol {
		println("You were overwritten by the hunter")
	}
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
		if maze[player.X][player.Y] == exit {
			// TODO: add something better to the win
			println("YOU WIN!!")
			winState = 1
			return maze
			// syscall.Exit(0)
		}
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
		player.Crumbs = make([]Coords, 20)
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

// renders with the information that the player has access to
func (player *Player) render(maze [][]string) {
	var infoRow []string

	infoRow = append(infoRow, "Hunter")

}
