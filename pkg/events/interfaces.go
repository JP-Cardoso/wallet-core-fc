package events

import "time"

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{} //entrada de dados genéricos
	SetPayload(payload interface{})
}

type EventHandlerInterface interface {
	//Para executar um comando, ele precisa do evento
	Handle(event EventInterface)
}

/*
evento - evento criado
dispatcher -> dispatch(evento)
*/
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}
