package main

import (
	"encoding/json"
	"os"
)

type TotalJson struct {
  players Player `json:"players"`
	hunters Hunter `json:"hunters"`
	maze [][]string `json:"maze"`
}

// side effect, modifies the global variable enemy. this makes sense as enemy is not effected by the runtime of this program but
// the program is effected by enemy
func syncInfo(maze [][]string, player Player) ([][]string, Player) {
	data, err := os.ReadFile("crossInfo.json")
	if err != nil {
		panic("could not read synced file")
	}
	var jsonData TotalJson
	err = json.Unmarshal(data, jsonData)
	if err != nil {
		println(err.Error())
		panic("error unmarshalling json")
	}
	enemy = jsonData.hunters

	return nil, *newPlayer()
}