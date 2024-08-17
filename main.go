package main

import (
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
  // var mazeSize int
  // fmt
	//   generateMaze(5)
	edgeList, tempMaze := generateMazeStruct(50)
	edgeList, size := krusals(edgeList, tempMaze)
	maze := drawMaze(edgeList, size)
	player := Player{X: 1, Y: 1, Symbol: TermBlue + " R " + TermReset}
	maze = player.placePlayer(maze)
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
		render(player.viewPort(maze, viewPortSize))
		turns++
		if turns%moveStep == 0 {
			userInput, err = keyboard.GetKeys(moveStep)
			if err != nil {
				panic(err)
			}
		}
    time.Sleep(time.Second / 100)
	}
}
