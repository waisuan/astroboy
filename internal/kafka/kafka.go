package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"net"
	"strconv"
)

const TOPIC = "my-topic"

type KafkaCli struct {
	addr string
}

func NewKafkaCli() *KafkaCli {
	return &KafkaCli{
		addr: "localhost:29092",
	}
}

func (k *KafkaCli) CreateTopic() error {
	conn, err := kafka.Dial("tcp", k.addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return err
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             TOPIC,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		return err
	}

	return nil
}

func (k *KafkaCli) ConsumeMessage(mailbox chan string) error {
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{k.addr},
		Topic:     TOPIC,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	//r.SetOffset(42)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(fmt.Errorf("failed to read message: %e", err))
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

		mailbox <- string(m.Value)
	}

	if err := r.Close(); err != nil {
		return err
	}

	return nil
}

func (k *KafkaCli) ProduceMessage(msgs kafka.Message) error {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:     kafka.TCP(k.addr),
		Topic:    TOPIC,
		Balancer: &kafka.LeastBytes{},
	}

	//kafka.Message{
	//	Key:   []byte("Key-A"),
	//	Value: []byte("Hello World!"),
	//},
	//	kafka.Message{
	//		Key:   []byte("Key-B"),
	//		Value: []byte("One!"),
	//	},
	//	kafka.Message{
	//		Key:   []byte("Key-C"),
	//		Value: []byte("Two!"),
	//	},
	err := w.WriteMessages(context.Background(), msgs)
	if err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}
