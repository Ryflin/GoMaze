package main

import (

)

func main() {
//   generateMaze(5)
  edgeList, tempMaze := generateMazeStruct(10)
  edgeList, size := krusals(edgeList, tempMaze)
  maze := drawMaze(edgeList, size)
//   drawMaze([]Edge{{X: 0, Y: 0, Dir: 2}, {X: 1, Y: 0, Dir: 0}}, 3)

  // game loop
  for ; ; {
    
  }
}
