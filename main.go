package main

import (
	"github.com/bpina/neptune/events"
	"github.com/bpina/neptune/game"
)

func main() {
	ps := new(events.ConsoleSubscriber)

	eb := events.NewEventBroadcaster()
	eb.AddSubscriber(ps)

	g := game.NewGame(eb)
	g.Initialize()
}
