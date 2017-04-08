package main

import "github.com/bpina/neptune/events"

func main() {
	ps := new(events.ConsoleSubscriber)

	eb := events.NewEventBroadcaster()
	eb.AddSubscriber(ps)

	e := events.NewEvent("test-event")
	e.Data["event-time"] = "now"

	eb.Broadcast(e)
}
