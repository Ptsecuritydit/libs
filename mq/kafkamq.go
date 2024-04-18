package mq

import (
	"github.com/Ptsecuritydit/libs/configs"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

type KafkaMq struct {
	Producer *kafka.Producer
	consumer *kafka.Consumer
}

func NewKafkaProducer(kafkaConf configs.KafkaConfig) KafkaMq {

	producer, err := kafka.NewProducer(configure(kafkaConf))

	if err != nil && configs.ServiceConfig.UseKafka {
		log.Fatal("Fail create KafkaConfig producer ")
	}

	return KafkaMq{Producer: producer}
}

func configure(config configs.KafkaConfig) *kafka.ConfigMap {

	kafkaConfig := kafka.ConfigMap{}

	for key, value := range config.Config {
		_ = kafkaConfig.SetKey(key, value)
	}
	return &kafkaConfig
}

// SendMassages sync
// send kafka message to internal queue. waiting
func (kafkaMq *KafkaMq) SendMassages(message kafka.Message) {

	delChan := make(chan kafka.Event)

	err := kafkaMq.Producer.Produce(&message, delChan)

	if err != nil {
		log.Println(err.Error())
	}
	answer := <-delChan
	msg := answer.(*kafka.Message)

	if msg.TopicPartition.Error != nil {
		log.Println(msg.TopicPartition.Error.Error())
	}

	close(delChan)
}

// SendMassageAsync
// send kafka message to internal queue.
// We are expecting delivery
func (kafkaMq *KafkaMq) SendMassageAsync(message kafka.Message) {
	err := kafkaMq.Producer.Produce(&message, nil)
	if err != nil {
		log.Println(err.Error())
	}

	go func() {
		for e := range kafkaMq.Producer.Events() {
			switch msg := e.(type) {
			case *kafka.Message:
				if msg.TopicPartition.Error != nil {
					log.Println(msg.TopicPartition.Error.Error())
				}
			}
		}
	}()
}
