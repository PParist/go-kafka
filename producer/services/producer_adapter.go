package services

import (
	"encoding/json"
	"events"
	"fmt"
	"reflect"

	"github.com/IBM/sarama"
)

type eventProducer struct {
	producer sarama.SyncProducer
}

func NewEventProducer(producer sarama.SyncProducer) EventProducer {
	return &eventProducer{producer: producer}
}

func (s *eventProducer) Produce(event events.Event) error {
	topic := reflect.TypeOf(event).Name()
	value, err := json.Marshal(event)
	if err != nil {
		return err
	}
	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}

	partition, offset, err := s.producer.SendMessage(&msg)
	if err != nil {
		return err
	}

	fmt.Println("Partition is", partition)
	fmt.Println("Offset is", offset)

	return nil
}
