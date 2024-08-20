package main

import "strings"

type node struct {
	coord    Coords
	contains string
	visited bool 
	steps int
}
type dikNode struct {
	visted bool
	steps  int
	node   node
}
func isNotWall(piece string) bool {
 return !strings.ContainsAny(piece, " ")
}


// func createEdgeList(maze [][]string) (edges []Edge) {
// 	for i, row := range maze {
// 		// check each square, then add each four directions.
// 		for j, val := range row {
// 			if isNotWall(val) {
// 				for dir:= 0; dir < 4; dir ++ {
// 					newX := 
// 				}
// 			}
// 		}
// 	}
// }

// func dikstras(maze [][]string, hunter Hunter) (nextMove Coords) {
// 	// steps to solve this problem. create another matrix that stores the following info
// 	// visited. length to visit. then to find the path I go backwards on the smallest value

// 	for {
	
// 	}
