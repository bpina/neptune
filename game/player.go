package game

import "fmt"

type Player struct {
	Name string
	Game *Game
}

func NewPlayer(name string, game *Game) *Player {
	p := new(Player)
	p.Name = name

	go p.receiveGameActions()

	return p
}

func (p *Player) receiveGameActions() {
	for {
		select {
		case action := <-p.Game.GameActions:
			p.HandleAction(action)
		case <-p.Game.Quit:
			return
		}
	}
}

func (p *Player) sendPlayerAction(name string) {
	a := Action{Name: name}

	p.Game.PlayerActions <- a
}

func (p *Player) HandleAction(a Action) {
	fmt.Printf("Received Action: %s\n", a.Name)
}
