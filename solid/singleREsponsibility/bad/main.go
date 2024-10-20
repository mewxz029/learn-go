package main

import "fmt"

type Player struct{}

func (p *Player) Move() {
	fmt.Println("Move")
}

func (p *Player) Attack() {
	fmt.Println("Attack")
}

func (p *Player) DisplayScoreBoard() {
	fmt.Println("DisplayScoreBoard")
}

func main() {
	p := &Player{}

	p.Move()
	p.Attack()
	p.DisplayScoreBoard()
}
