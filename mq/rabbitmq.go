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

func (r *RabbitMQ) RabbitQueueDeclare(qname string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) error {
	_, err := r.ch.QueueDeclare(qname, durable, autoDelete, exclusive, noWait, args)
	return err
}

func (r *RabbitMQ) ChannelPublish(exchange string, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	return r.ch.Publish(exchange, key, mandatory, immediate, msg)
}
