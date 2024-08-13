package main

type Player struct {
	X      int
	Y      int
	Hunter bool
	Symbol string
}

// move the player around
func (this *Player) move(maze [][]string, dir int) {
	newX := this.X + dirs[dir][0]
	newY := this.Y + dirs[dir][1]
	if maze[newX][newY] != wallRep {
		maze[this.X][this.Y] = sqrRep
		this.X = newX + 1
		this.Y = newY + 1
		maze[this.X][this.Y] = this.Symbol
	}
}
