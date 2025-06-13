package producer_test

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"net"
	"strconv"
	"time"
)

var Topic = "test"
var Partition = 0
var GroupID = "my_group1"

func ProducerTest() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9094", Topic, Partition)
	if err != nil {
		log.Fatal("连接leader失败:", err)
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("一!")},
		kafka.Message{Value: []byte("二!")},
		kafka.Message{Value: []byte("三!")},
	)
	if err != nil {
		log.Fatal("写入消息失败:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("关闭写入器失败:", err)
	}
}

func ProducerTest2() {
	// 通过现有的非leader连接连接到Kafka leader，而不是使用DialLeader
	conn, err := kafka.Dial("tcp", "localhost:9094")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var connLeader *kafka.Conn
	connLeader, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer connLeader.Close()
}
func ProducerTest3() {
	// 创建一个新的读取器，从主题A的分区0的偏移量42开始消费
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		Topic:     "topic-A",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(42)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("偏移量 %d 的消息: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("关闭读取器失败:", err)
	}
}
