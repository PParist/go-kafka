package services

import (
	"github.com/IBM/sarama"
)

type consumerHandler struct {
	eventHandler EventHandler
}

func NewConsumerHandler(eventHandler EventHandler) sarama.ConsumerGroupHandler {
	return &consumerHandler{eventHandler: eventHandler}
}

func (s *consumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}
func (s *consumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
func (s *consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	//TODO: for message
	for msg := range claim.Messages() {
		//Mark for ignor already read msg
		s.eventHandler.Handle(msg.Topic, msg.Value)
		session.MarkMessage(msg, "")
	}
	return nil
}
