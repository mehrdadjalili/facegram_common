package log

import (
	"github.com/mehrdadjalili/facegram_common/pkg/logger/models"
	"github.com/streadway/amqp"
)

type (
	option struct {
		broker  *amqp.Connection
		service models.Service
	}
	Log interface {
		Info()
		Warning()
		Error()
	}
)

func New(url string, service models.Service) (Log, error) {
	co, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return &option{broker: co, service: service}, nil
}

func (o *option) Info() {
	panic("implement me")
}

func (o *option) Warning() {
	panic("implement me")
}

func (o *option) Error() {
	panic("implement me")
}
