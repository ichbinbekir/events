package events

type EventEmitter struct {
	events map[string]chan []any
}

func NewEventEmitter() *EventEmitter {
	eventemitter := new(EventEmitter)
	eventemitter.events = map[string]chan []any{}
	return eventemitter
}

func (eventemitter *EventEmitter) On(eventName string, listener func(...any)) {
	eventemitter.events[eventName] = make(chan []any)
	go func() {
		for {
			args := <-eventemitter.events[eventName]
			listener(args...)
		}
	}()
}

func (eventemitter *EventEmitter) Once(eventName string, listener func(...any)) {
	eventemitter.events[eventName] = make(chan []any)
	go func() {
		args := <-eventemitter.events[eventName]
		listener(args...)
	}()
}

func (eventemitter *EventEmitter) Emit(eventName string, args ...any) {
	if eventemitter.events[eventName] != nil {
		eventemitter.events[eventName] <- args
	}
}

func (eventemitter *EventEmitter) Off(eventName string) {
	eventemitter.events[eventName] = nil
}

func (eventemitter *EventEmitter) OffAll() {
	eventemitter.events = nil
}
