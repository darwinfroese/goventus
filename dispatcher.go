package main

import "fmt"

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
func notify(t EventType, e Event) int {
	respondents := 0

	for _, h := range eventMap[t] {
		h(e)
		respondents++
	}

	return respondents
}

// Run is a coroutine for receiving events and notifing observers
func Run(notifier chan EventObject) {
	for {
		e := <-notifier

		fmt.Println("DEBUG - Event Received")

		notify(e.EventType, e.Event)
	}
}
