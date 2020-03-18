package main

import (
	"context"
	"log"
	"os"
	"fmt"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

const (
	filename = "out.txt"
	// KAFKA_URL = "localhost:9092"
	// KAFKA_TOPIC = "streams-wordcount-processor-output"
	// KAFKA_GROUP_ID = "TestGroupID"
	KAFKA_URL = "KAFKA_URL"
	KAFKA_TOPIC = "KAFKA_TOPIC"
	KAFKA_GROUP_ID = "KAFKA_GROUP_ID"
)


func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		// GroupID:  groupID,
		Topic:    topic,
		Partition: 0,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func main() {
	// get kafka reader using environment variables.
	kafkaURL, ok := os.LookupEnv(KAFKA_URL)
	if !ok {
		fmt.Println("Error getting KAFKA URL from environment variable")
	}

	topic, ok := os.LookupEnv(KAFKA_TOPIC)
	if !ok {
		fmt.Println("Error getting KAFKA TOPIC from environment variable")
	}

	groupID, ok := os.LookupEnv(KAFKA_GROUP_ID)
	if !ok {
		fmt.Println("Error getting KAFKA GROUP ID from environment variable")
	}

	reader := getKafkaReader(kafkaURL, topic, groupID)

	defer reader.Close()

	tbl := map[string]string{}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error Opening output file", err)
		return
	}
	
	defer f.Close()

	// AWSClean()
	// AWSInit()
	// defer AWSClean()

	fmt.Println("start consuming stream ...")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading stream:", err)
			os.Exit(1)
		}

		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		tbl[string(m.Key)] = string(m.Value)

		err = os.Truncate(filename, 0)
		if err != nil {
			fmt.Println("Error Truncate ->", err)
		}

	
		for key, value := range tbl {
			fmt.Println(key, ":", value)
			fmt.Fprintf(f, "%s : %s\n", string(key), string(value))
			// AWSUpload(abucket, filename)
			GelatoUpload("bkt002", filename)
		}
	}

	// AWSClean()

}
