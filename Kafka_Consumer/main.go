package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

func initKafka() *kgo.Client {
	seeds := []string{"localhost:9093"}

	cl, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ConsumerGroup("my-test-group"),
		kgo.ConsumeTopics("test-topic"),

		// for development only
		kgo.AllowAutoTopicCreation(),
	)
	if err != nil {
		log.Fatalf("unable to create kafka client: %v", err)
	}
	return cl
}
func consumeFromTopic(cl *kgo.Client) {
	ctx := context.Background()
	consumeCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	fmt.Println("⏳ Waiting for messages...")

	fetches := cl.PollFetches(consumeCtx)
	if errs := fetches.Errors(); len(errs) > 0 {
		for _, e := range errs {
			log.Printf("fetch error: topic=%s partition=%d err=%v", e.Topic, e.Partition, e.Err)
		}
		return
	}

	iter := fetches.RecordIter()
	for !iter.Done() {
		rec := iter.Next()
		fmt.Printf("📥 Consumed: topic=%s partition=%d offset=%d key=%s value=%s\n",
			rec.Topic, rec.Partition, rec.Offset, string(rec.Key), string(rec.Value))
	}
}

func main() {
	kafkaClient := initKafka()
	defer kafkaClient.Close()

	for {
		consumeFromTopic(kafkaClient)
	}
}
