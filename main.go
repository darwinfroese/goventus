package main

import "fmt"

// Version is the current version of Goventus
var (
	version string
)

func main() {
	fmt.Println("Goventus", version)

	notifier := make(chan EventObject)
	go Run(notifier)

	Subscribe(LogNothing, Log)
	notifier <- EventObject{
		EventType: Log,
		Event:     LogEvent{Message: "Hello World"},
	}

	for {
	}
}

// LogEvent is an event type for logging
type LogEvent struct {
	Message string
}

// LogNothing literally logs nothing
func LogNothing(e Event) {
	if e, ok := e.(LogEvent); ok {
		fmt.Println(e.Message)
	}
}
