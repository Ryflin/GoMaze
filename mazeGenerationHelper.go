package main

import (
	"encoding/json"
	"fmt"
)

func printJson(thing any) {
	thingBytes, err := json.Marshal(thing)
	if err != nil {
		thingBytes = []byte{}
	}
	fmt.Println(string(thingBytes))
}

func oneNumber(maze [][]int) bool {
	var number = maze[0][0]
	for _, row := range maze {
		for _, val := range row {
			if val != number {
				return false
			}
		}
	}
	return true
}

func mapDoesNotContain(maze map[int][][]int, x int, y int, val int) bool {
	array, exists := maze[val]
	if !exists {
		return false
	}
	println(len(array))
	for i := 0; i < len(array); i++ {
		if array[i][0] == y && array[i][1] == x {
			return true
		}
	}
	return false
}

// add all the values in s1 to s2
func addSets(maze map[int][][]int, set1 int, set2 int) map[int][][]int {
	ar1 := maze[set1]
	ar2 := maze[set2]
	ar2 = append(ar2, ar1...)
	delete(maze, set1)
	maze[set2] = ar2
	return maze
}

// removes the listed element.
//
// requires that that element exists in the edge list
func findAndRemove(a []Edge, e Edge) []Edge {
	for i, ed := range a {
		if ed == e {
			copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
			a[len(a)-1] = Edge{}    // Erase last element (write zero value).
			a = a[:len(a)-1]     // Truncate slice.
			// a[i] = a[len(a)-1] // Copy last element to index i.
			// a[len(a)-1] = ""   // Erase last element (write zero value).
			// a = a[:len(a)-1]   // Truncate slice.
			break
		}
	}
	return a
}

// all val1's become val2's
func swapAll(maze [][]int, val1 int, val2 int) [][]int {
	for i, row := range maze {
		for j, val := range row {
			if val == val1 {
				maze[i][j] = val2
			}
		}
	}
	return maze
}
