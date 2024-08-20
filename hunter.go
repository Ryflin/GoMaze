package main

type Hunter struct {
	X int `json:"X"`
	Y int	`json:"Y"`
	Symbol string
	npcDifficulty int
}

// func (hunter *Hunter) move(maze [][]string) ([][]string){

// }

// // moves the hunter as if it is an npc
// func (hunter *Hunter)moveNpc(maze [][]string) ([][]string) {
// // I am finding that this is producing a lot of copies of map, I don't like that but progamming it differently would cause side effects. 

// }