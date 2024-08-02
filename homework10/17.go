package main

type Event struct {
	name string
	date string
}
type EventList struct {
	Events []Event
}

func (e *EventList) AddEvent(name, date string) {
	e.Events = append(e.Events, Event{name, date})
}