package main

import "fmt"

type (
	Player     struct{}
	ScoreBoard struct{}
)

func (p *Player) Move() {
	fmt.Println("Move")
}

func (p *Player) Attack() {
	fmt.Println("Attack")
}

func (s *ScoreBoard) DisplayScoreBoard() {
	fmt.Println("DisplayScoreBoard")
}

func main() {
	p := &Player{}
	sb := &ScoreBoard{}

	p.Move()
	p.Attack()

	sb.DisplayScoreBoard()
}
