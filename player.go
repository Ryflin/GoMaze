package main


type Player struct {
	X      int
	Y      int
	Hunter bool
	Symbol string
}

// move the player around
func (this *Player) move(maze [][]string, dir int) (newMaze [][]string) {
	newX := this.X + dirs[dir][0]
	newY := this.Y + dirs[dir][1]
	if maze[newX][newY] == emptyTile {
		mazeMutex.Lock()
		maze[this.X][this.Y] = emptyTile
		this.X = newX + dirs[dir][0]
		this.Y = newY + dirs[dir][1]
		maze[this.X][this.Y] = this.Symbol
		mazeMutex.Unlock()
	}
	return maze
}

func (this *Player) placePlayer(maze [][]string) (newMaze [][]string) {
	for maze[this.X][this.Y] != emptyTile {
		if maze[this.X][this.Y+1] == emptyTile {
			this.Y += 1
		} else {
			this.X += 1
		}
	}
	mazeMutex.Lock()
	maze[this.X][this.Y] = this.Symbol
	mazeMutex.Unlock()
	return maze
}
