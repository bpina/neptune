package game

import (
	"fmt"
	"github.com/bpina/neptune/events"
)

type Game struct {
	Players       []*Player
	PlayerCount   int
	Broadcaster   *events.EventBroadcaster
	GameActions   chan Action
	PlayerActions chan Action
	Quit          chan int
}

type Action struct {
	Name string
}

func NewGame(eb *events.EventBroadcaster) *Game {
	g := new(Game)
	g.Broadcaster = eb
	g.GameActions = make(chan Action)
	g.PlayerActions = make(chan Action)
	g.Quit = make(chan int)

	return g
}

func (g *Game) AddPlayer(p *Player) {
	p.Game = g

	if g.Players != nil {
		g.Players = append(g.Players, p)
	} else {
		g.Players = []*Player{p}
	}

	e := events.NewEvent("player-added")
	e.Data["player-name"] = p.Name

	g.BroadcastEvent(e)
}

func (g *Game) BroadcastEvent(e *events.Event) {
	g.Broadcaster.Broadcast(e)
}

func (g *Game) Initialize() {
	p1 := NewPlayer("Player 1", g)
	g.AddPlayer(p1)

	p2 := NewPlayer("Player 2", g)
	g.AddPlayer(p2)

	g.PlayerCount = 2

	e := events.NewEvent("game-initialized")
	g.BroadcastEvent(e)

	g.sendAction("Phase Begin")
}

func (g *Game) sendAction(name string) {
	a := Action{Name: name}

	g.GameActions <- a
}

func (g *Game) receivePlayerActions() {
	for {
		select {
		case action := <-g.PlayerActions:
			g.HandlePlayerAction(action)
		case <-g.Quit:
			return
		}
	}
}

func (g *Game) End() {
	g.Quit <- 0
}

func (p *Game) HandlePlayerAction(a Action) {
	fmt.Printf("Received Action: %s\n", a.Name)
}
