package main

// EventType is an event type alias
type EventType int

// Event is an event alias for arguments
type Event interface{}

// EventObject is a wrapper for event and eventtype
type EventObject struct {
	EventType
	Event
}

// Handler type for event handling, takes in any arguments
type Handler func(args Event)

// Event enumerator
const (
	Log EventType = iota
)
