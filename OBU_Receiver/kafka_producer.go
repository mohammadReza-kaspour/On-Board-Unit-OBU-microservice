package main

import (
	"context"
	"fmt"
	"log"

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
func produceToTopic(cl *kgo.Client, record *kgo.Record) {

	ctx := context.Background()

	cl.Produce(ctx, record, func(r *kgo.Record, err error) {
		if err != nil {
			log.Printf("failed to produce message: %v", err)
			return
		}
		fmt.Printf("Produced message to topic %s partition %d offset %d\n", r.Topic, r.Partition, r.Offset)
	})
}
