package core

import "reflect"

type Emitter struct {
	listeners map[string][]func()
}

func (e *Emitter) Init() {
	e.listeners = map[string][]func(){}
}

func (e *Emitter) On(event string, handler func()) {
	handlers, ok := e.listeners[event]
	if ok {
		e.listeners[event] = append(handlers, handler)
	} else {
		e.listeners[event] = []func(){handler}
	}
}

// func (e *Emitter) Once(event string, handler func())

func (e *Emitter) Off(event string, handler func()) {
	handlers, ok := e.listeners[event]
	if ok {
		for i, v := range handlers {
			if reflect.ValueOf(v).Pointer() == reflect.ValueOf(handler).Pointer() {
				e.listeners[event] = append(handlers[:i], handlers[i+1])
				return
			}
		}
	}
}

func (e *Emitter) Has(event string) bool {
	_, ok := e.listeners[event]
	return ok
}

func (e *Emitter) Emit(event string) {
	handlers, ok := e.listeners[event]
	if ok {
		for _, v := range handlers {
			v()
		}
	}
}

func (e *Emitter) Clear() {
	e.Init()
}
