package events

import "errors"

var (
	ErrHandlerAlreadyRegistered = errors.New("handler already registered")
	ErrHandlerNotRegistered     = errors.New("handler not registered")
)

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {

	// verifico se já existir um handler
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Has(eventName string, eventHandler EventHandlerInterface) bool {
	if list, ok := ed.handlers[eventName]; ok {
		for _, value := range list {
			if value == eventHandler {
				return true
			}
		}
	}

	return false
}

func (ed *EventDispatcher) Remove(eventName string, eventHandler EventHandlerInterface) error {
	if !ed.Has(eventName, eventHandler) {
		return ErrHandlerNotRegistered
	}

	// remover do slite

	return nil
}