package services

import (
	"events"
)

type EventProducer interface {
	Produce(event events.Event) error
}
