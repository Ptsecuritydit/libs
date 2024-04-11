package mq

import "github.com/streadway/amqp"

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQ(uri string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &RabbitMQ{conn, ch}, nil
}

func (r *RabbitMQ) Close() {
	r.conn.Close()
}

func (r *RabbitMQ) RabbitQueueDeclare(qname string) error {
	_, err := r.ch.QueueDeclare(qname, false, false, false, false, nil)
	return err
}

func (r *RabbitMQ) ChannelPublish(exchange string, key string, mendatory, immadiate bool, msg amqp.Publishing) error {
	return r.ch.Publish(exchange, key, mendatory, immadiate, msg)
}
