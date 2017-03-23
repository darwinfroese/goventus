package main

// EventType is an event type alias
type EventType int

// Event is an event alias for arguments
type Event interface{}

// Handler type for event handling, takes in any arguments
type Handler func(args Event)

// Event enumerator
const (
	Log EventType = iota
)

var eventMap map[EventType][]Handler
var eventCount = 0

func init() {
	eventMap = make(map[EventType][]Handler)
}

// Subscribe to a function
func Subscribe(h Handler, e EventType) int {
	if len(eventMap[e]) == 0 {
		eventMap[e] = make([]Handler, 0)
	}
	eventMap[e] = append(eventMap[e], h)

	eventCount++
	return eventCount
}

// Notify all subscribers to an event
func Notify(t EventType, e Event) int {
	respondents := 0

	for _, h := range eventMap[t] {
		h(e)
		respondents++
	}

	return respondents
}
