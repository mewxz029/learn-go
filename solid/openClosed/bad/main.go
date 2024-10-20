package main

type (
	Player struct{}
	Steak  struct{}
)

func (p *Player) EatSteak(steak *Steak) {
	steak.Heal()
	steak.DamageBluff()
}

func (s *Steak) Heal() {
	println("Heal")
}

func (s *Steak) DamageBluff() {
	println("Increase damage 100%")
}

func main() {
	p := &Player{}
	s := &Steak{}

	p.EatSteak(s)
}
