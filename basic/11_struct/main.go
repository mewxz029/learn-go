package main

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	Username string `json:"username"`
	Level    uint   `json:"level"`
	Status   string `json:"status"`
	Class    string `json:"class"`
}

func (p *Player) GetUsername() string {
	return p.Username
}

func (p *Player) LevelUp() {
	p.Level++
}

func main() {
	player1 := Player{
		Username: "Mew",
		Level:    1,
		Status:   "active",
		Class:    "Warrior",
	}

	player1.LevelUp()

	jsonData, _ := json.MarshalIndent(&player1, "", "\t")
	fmt.Println(string(jsonData))
	fmt.Println(player1.GetUsername())
}
