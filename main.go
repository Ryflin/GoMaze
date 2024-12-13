package main

import (
	"syscall"
	"time"

	"github.com/eiannone/keyboard"
)

// I wanted to add a cpu hunter that chases the player I am not concerned about client side security and such for now

func main() {
	// var mazeSize int
	// fmt
	//   generateMaze(5)

	edgeList, tempMaze := generateMazeStruct(20)
	edgeList, size := krusals(edgeList, tempMaze)
	maze := drawMaze(&edgeList, size)
	player := *newPlayer(withColor(CrumbBlue), withCoords(Coords{X: 1, Y: 1}), withSymbol(TermBlue+" R "+TermReset), withCrumbs(make([]Coords, 40)))
	maze = player.placePlayer(maze)
	playerView := player.viewPort(&maze, viewPortSize)
	makeNewUi(playerView)
	// TODO implement capture input (for now wars (wasd for colemak))
	// method capture key and re-render
	// game loop
	turns := 0
	userInput, err := keyboard.GetKeys(moveStep)
	if err != nil {
		println(err.Error())
		panic("I really don't know what to do here")
	}
	defer func() {
		keyboard.Close()
	}()
	time.Sleep(1000)
	for {
		event := <-userInput
		if event.Err != nil {
			println("error in the event")
			panic(event.Err)
		}
		if event.Key == keyboard.KeyCtrlC {
			panic("keyboard interupt")
		}
		if dir, exists := inputs[string(event.Rune)]; exists {
			maze = player.move(maze, dir, turns)
		}
		updateUI(player.viewPort(&maze, viewPortSize))
		// maze, player = syncInfo(maze, player)
		turns++
		if turns%moveStep == 0 {
			userInput, err = keyboard.GetKeys(moveStep)
			if err != nil {
				panic(err)
			}
		}
		time.Sleep(time.Second / 100)
		if winState == 1 {
			win(player)
		}
		if winState == 2 {
			lose(player)
		}
	}
}

func win(player Player) {
	println("you win good job")
	syscall.Exit(0)
}
func lose(player Player) {
	println("You lost sorry")
	syscall.Exit(0)
}
