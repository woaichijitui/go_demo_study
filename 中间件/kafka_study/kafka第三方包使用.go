package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const TOPIC = "test"

var reader *kafka.Reader

func WriteKafka(ctx context.Context) {
	writer := kafka.Writer{
		Addr:                   kafka.TCP("172.31.191.51:9092", "172.31.191.51:9095", "172.31.176.1:9097"),
		Topic:                  TOPIC,
		Balancer:               &kafka.Hash{},
		RequiredAcks:           kafka.RequireNone,
		WriteTimeout:           time.Second,
		AllowAutoTopicCreation: true,
	}
	defer writer.Close()

	for i := 0; i < 2; i++ {
		if err := writer.WriteMessages(
			ctx,
			kafka.Message{Key: []byte("11"), Value: []byte("你好1")},
			kafka.Message{Key: []byte("22"), Value: []byte("你好2")},
			kafka.Message{Key: []byte("33"), Value: []byte("你好3")},
			kafka.Message{Key: []byte("11"), Value: []byte("我是key ：1")},
			kafka.Message{Key: []byte("22"), Value: []byte("我是key ：2")},
			kafka.Message{Key: []byte("11"), Value: []byte("lalalala")},
		); err == nil {
			fmt.Println("写入成功")
			break
		} else if err != nil {

			if err == kafka.LeaderNotAvailable {
				time.Sleep(time.Second)
				continue
			}
			fmt.Println("写入失败", err)
		}
	}

}
func ReadKafka(ctx context.Context) {
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"172.31.176.1:9092", "172.31.176.1:9095", "172.31.176.1:9097"},
		Topic:          TOPIC,
		CommitInterval: 1 * time.Second,
		GroupID:        "my_group1",
		StartOffset:    kafka.FirstOffset,
	})
	//defer reader.Close()

	for {
		message, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Println("读取kafka失败")
			break
		}

		fmt.Printf("topic: %v, groupId: %v, patition: %v, offset: %v,key: %v,message: %v\n", message.Topic, message.Headers, message.Partition, message.Offset, string(message.Key), string(message.Value))
	}

}

func ListenSingal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	fmt.Println("收到信号：", sig)
	reader.Close()
	os.Exit(0)
}

func test() {
	ctx := context.Background()
	WriteKafka(ctx)
	go ListenSingal()
	ReadKafka(ctx)
}
