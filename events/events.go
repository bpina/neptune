package events

import "fmt"

type EventBroadcaster struct {
	Messages    chan Event
	Subscribers []EventSubscriber
}

type EventSubscriber interface {
	EventReceived(*Event)
}

type Event struct {
	Key       string
	Timestamp int
	Data      map[string]string
}

func NewEvent(key string) *Event {
	e := new(Event)
	e.Key = key
	e.Data = make(map[string]string)

	return e
}

func NewEventBroadcaster() *EventBroadcaster {
	eb := new(EventBroadcaster)
	eb.Messages = make(chan Event)

	return eb
}

func (eb *EventBroadcaster) AddSubscriber(es EventSubscriber) {
	if eb.Subscribers == nil {
		eb.Subscribers = []EventSubscriber{es}
	} else {
		eb.Subscribers = append(eb.Subscribers, es)
	}
}

func (eb *EventBroadcaster) Broadcast(e *Event) {
	if eb.Subscribers != nil {
		for _, sub := range eb.Subscribers {
			sub.EventReceived(e)
		}
	}
}

type ConsoleSubscriber struct{}

func (ps *ConsoleSubscriber) EventReceived(e *Event) {
	fmt.Printf("Event: %s\n", e.Key)

	if e.Data != nil {
		for k, v := range e.Data {
			fmt.Printf("Data: %s\nValue: %s\n", k, v)
		}
	}
}
