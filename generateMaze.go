package main

import (
	"fmt"
	"math/rand"
)

type Edge struct {
	X   int
	Y   int
	Dir int
}

// each node is part of a set. that set then is checked
type Node struct {
	Data int
	X    int
	Y    int
}

// generates a list of edges and a 2d array that is the graph
func generateMazeStruct(size int) (edgeList []Edge, maze [][]int) {
	maze = make([][]int, size)
	for i := 0; i < size; i++ {
		maze[i] = make([]int, size)
	}
	// now generate edges
	index := 0
	for i, row := range maze {
		for j := range row {
			maze[i][j] = index
			index += 1
			if i != 0 {
				newEdge := Edge{
					X:   i,
					Y:   j,
					Dir: 0,
				}
				edgeList = append(edgeList, newEdge)
			}
			if j != 0 {
				newEdge := Edge{
					X:   i,
					Y:   j,
					Dir: 1,
				}
				edgeList = append(edgeList, newEdge)

			}
			if i != len(maze)-1 {
				newEdge := Edge{
					X:   i,
					Y:   j,
					Dir: 2,
				}
				edgeList = append(edgeList, newEdge)

			}
			if j != len(row)-1 {
				newEdge := Edge{
					X:   i,
					Y:   j,
					Dir: 3,
				}
				edgeList = append(edgeList, newEdge)
			}

		}
	}
	// in theory edge list should now be every edge
	// fmt.Println(edgeList)
	// now start kursals algorithm
	return edgeList, maze
}

// krusal's algorithm
//
// run once or twice no need to optimize (yet)
func krusals(list []Edge, maze [][]int) (edgeList []Edge, size int) {
	rand.Shuffle(len(list), func(i, j int) {
		tempEdge := list[i]
		list[i] = list[j]
		list[j] = tempEdge
	})
	var index = 0
	// after being removed then also remove the backwards implementation
	for !oneNumber(maze) {
		// printJson(edgeList)
		// printJson(maze)
		// time.Sleep(time.Second)
		// removes the first element and the element that matches that element
		tempEdge := list[index]
		var nextNode = []int{tempEdge.X + dirs[tempEdge.Dir][0], tempEdge.Y + dirs[tempEdge.Dir][1]}
		// if out of bounds just try again this should in theory never trigger as the bounds checking is done when making the edge
		if nextNode[0] < 0 || nextNode[0] >= len(maze) || nextNode[1] < 0 || nextNode[1] >= len(maze) {
			index += 1
			fmt.Println("you messed up. index is out of bounds, check the edge maing function.")
			fmt.Println(nextNode)
			// time.Sleep(time.Second)
			continue
		}
		// fmt.Println(tempEdge, nextNode)
		if maze[tempEdge.X][tempEdge.Y] == maze[nextNode[0]][nextNode[1]] {
			index += 1
			continue
		}
		// this needs to be all and it is currently just one.
		swapAll(maze, maze[nextNode[0]][nextNode[1]], maze[tempEdge.X][tempEdge.Y])
		// maze[nextNode[0]][nextNode[1]] = maze[tempEdge.X][tempEdge.Y]
		findAndRemove(list, tempEdge)
		findAndRemove(list, Edge{
			X:   nextNode[0],
			Y:   nextNode[1],
			Dir: (tempEdge.Dir + 2) % 4,
		})
	}
	// remove null values in slice
	var tempList []Edge
	for _, edge := range list {
		empty := Edge{}
		if edge != empty {
			tempList = append(tempList, edge)
		}
	}
	// fmt.Println("final list: ", tempList)
	// TODO: place this into generate maze function for modularity and better logging
	return tempList, len(maze)
	// drawMaze(tempList, len(maze))
}

// draws the maze that is represented in the edge list.
//
// run once or twice, no need to optimize
func drawMaze(list []Edge, size int) (maze [][]string) {
	// every other is the edge so then add whatever is in dirs. then look at maze and see what the deal is
	maze = make([][]string, size*2-1)
	for i := 0; i < len(maze); i++ {
		// fmt.Print(" i:", i)
		maze[i] = make([]string, len(maze))
	}
	for i := range maze {
		for j := range maze[i] {
			maze[i][j] = sqrRep
		}
		if i%2 != 0 {
			for j := range maze[i] {
				maze[i][j] = "+"
				maze[j][i] = "+"
			}
		}
	}
	// new maze creation handled, now on to filling maze based on list
	for _, edge := range list {
		x := edge.X*2 + dirs[edge.Dir][0]
		y := edge.Y*2 + dirs[edge.Dir][1]
		maze[x][y] = wallRep
		if x+1 < len(maze) && x+1 > 0 && maze[x+1][y] == "+" {
			maze[x+1][y] = wallRep
		}
		if y+1 < len(maze) && y+1 > 0 && maze[x][y+1] == "+" {
			maze[x][y+1] = wallRep
		}
		if x-1 < len(maze) && x-1 > 0 && maze[x-1][y] == "+" {
			maze[x-1][y] = wallRep
		}
		if y-1 < len(maze) && y-1 > 0 && maze[x][y-1] == "+" {
			maze[x][y-1] = wallRep
		}
	}

	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] == "+" {
				maze[i][j] = sqrRep
			}
		}
	}
	var replaceMaze [][]string
	var tempMaze []string
	tempMaze = append(tempMaze, wallRep)
	for i := 0; i < len(maze); i++ {
		tempMaze = append(tempMaze, wallRep)

	}
	tempMaze = append(tempMaze, wallRep)
	replaceMaze = append(replaceMaze, tempMaze)
	for i := range maze {
		var tempMaze []string
		tempMaze = append(tempMaze, wallRep)
		tempMaze = append(tempMaze, maze[i]...)
		tempMaze = append(tempMaze, wallRep)
		// fmt.Println(tempMaze)
		replaceMaze = append(replaceMaze, tempMaze)
	}
	var tempMaze2 []string
	tempMaze2 = append(tempMaze2, wallRep)
	for i := 0; i < len(maze); i++ {
		tempMaze2 = append(tempMaze2, wallRep)

	}
	tempMaze2 = append(tempMaze2, wallRep)
	replaceMaze = append(replaceMaze, tempMaze2)
	// fmt.Println(replaceMaze)
	maze = replaceMaze
	// for now have a for loop at the bottom displaying all the stuff

	// printJson(maze)
	for _, row := range maze {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
	return maze
}
