package services

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}
