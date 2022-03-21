package adapter

import (
	"fmt"
	"log"
	"net"
	"time"
	"errors"
	"context"

	"github.com/streadway/amqp"
)

type AmqpClient struct {
	Conn *amqp.Connection
	Channel *amqp.Channel
}

func GetAmqpClient(amqpConnString string) *AmqpClient {
	conn, channel, _ := GetAmqpConnAndChannel(amqpConnString)

	return &AmqpClient{
		Conn: conn,
		Channel: channel,
	}
}

func GetAmqpConnAndChannel(amqpConnString string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.DialConfig(amqpConnString, amqp.Config{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, 30 * time.Second)
		},
	})
	if err != nil {
		log.Fatalf("Failed to connect to AMQP endpoint: %v", err)
		return nil, nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to connect to AMQP endpoint: %v", err)
		return nil, nil, err
	}

	return conn, channel, nil
}

func (client AmqpClient) BuildMessageJson(messageJson string, replyQueueName string) amqp.Publishing {
	messageBody := amqp.Publishing{
		ContentType: "application/json",
		Body: []byte(messageJson),
		ReplyTo: replyQueueName,
		DeliveryMode: amqp.Persistent,
	}
	if replyQueueName != "" {
		messageBody.ReplyTo = replyQueueName
	}
	return messageBody
}

func (client AmqpClient) PublishMessage(message amqp.Publishing, exchange string, routingKey string, priority uint8) error {
	messageBody := amqp.Publishing{
		DeliveryMode: message.DeliveryMode,
		Timestamp: time.Now(),
		ContentType: message.ContentType,
		ReplyTo: message.ReplyTo,
		Body: message.Body,
	}
	if priority != 0 {
		messageBody.Priority = priority
	}
	return client.Channel.Publish(
		exchange, // exchange
		routingKey, // queue name or routing key
		false, // mandatory
		false, // immediate
		messageBody,
	)
}

func (client AmqpClient) MakeTempQueue() (amqp.Queue, error) {
	return client.Channel.QueueDeclare(
		"", // name
		false, // durable
		true, // delete when unused
		true, // exclusive
		false, // no-wait
		nil, // arguments
	)
}

func (client AmqpClient) ConsumeMessageFromTempQueue(queueName string, consumerName string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var msgs (<-chan amqp.Delivery)
	var messageBody []byte

	msgs, err := client.Channel.Consume(
		queueName, // queue
		consumerName, // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil, // args
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to Consume from Queue %s: %v", queueName, err.Error()))
	}
	if msgs != nil {
		cancel()
	}

	select {
		case <-ctx.Done():
			if msgs == nil {
				return nil, errors.New(fmt.Sprintf("Failed to Consume from Queue %s: %v", queueName, ctx.Err()))
			}
	}

	for msg := range msgs {
		messageBody = []byte(msg.Body)
		msg.Ack(true)
		// cancel the consumer, then the queue will be deleted
		if err := client.Channel.Cancel(consumerName, false); err != nil {
			return nil, errors.New(fmt.Sprintf("Failed to cancel consumer %s: %v", consumerName, err.Error()))
		}
		break
	}

	return messageBody, nil
}
