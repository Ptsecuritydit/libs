package mq

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s\n", msg, err)
	}
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

func (r *RabbitMQ) Close() error {
	return r.conn.Close()
}

func (r *RabbitMQ) RabbitQueueDeclare(qname string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) error {
	_, err := r.ch.QueueDeclare(qname, durable, autoDelete, exclusive, noWait, args)
	return err
}

func (r *RabbitMQ) ChannelPublish(exchange string, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	return r.ch.Publish(exchange, key, mandatory, immediate, msg)
}

func (r *RabbitMQ) RabbitConsume(exchange string, key string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	return r.ch.Consume(exchange, key, autoAck, exclusive, noLocal, noWait, args)
}

func (r *RabbitMQ) RabbitExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error {
	return r.ch.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args)
}

func (r *RabbitMQ) RabbitQueueBind(name, key, exchange string, noWait bool, args amqp.Table) error {
	return r.ch.QueueBind(name, key, exchange, noWait, args)
}
