/**
@description	kafka 简单的消费者(consumer)，支持一个消费者消费多个topic，consumer-group和consumer负载均衡
	sarama-cluster，kafka go语言客户端库。支持kafka0.8及以上版本。
@github https://github.com/bsm/sarama-cluster
@author pujie
@data	2018-01-17
**/

package main

import (
	"fmt"
	"github.com/Shopify/sarama"

	"log"
	"os"
	"os/signal"
	"sync"
)

var Address = []string{"10.14.41.57:9092", "10.14.41.58:9092", "10.14.41.59:9092"}

func main() {
	topic := "d-index-serving-inverted-update"
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	// 广播式消费：消费者1
	go clusterConsumer(wg, Address, topic, "127.0.0.1")
	// 广播式消费：消费者2
	go clusterConsumer(wg, Address, topic, "10.130.151.41")

	wg.Wait()
}

// 支持brokers cluster的消费者
func clusterConsumer(wg *sync.WaitGroup, brokers []string, topic, groupId string) {
	defer wg.Done()
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// init consumer
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Printf("%s: sarama.NewSyncProducer err, message=%s \n", groupId, err)
		return
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partition, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partition.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// trap SIGINT to trigger a shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range partition.Errors() {
			log.Printf("%s:Error: %s\n", groupId, err.Error())
		}
	}()

	// consume messages, watch signals
	var successes int
Loop:
	for {
		select {
		case msg, ok := <-partition.Messages():
			if ok {
				fmt.Fprintf(os.Stdout, "%s:%s/%d/%d\t%s\t%s\n", groupId, msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				successes++
			}
		case <-signals:
			break Loop
		}
	}
	fmt.Fprintf(os.Stdout, "%s consume %d messages \n", groupId, successes)
}
