package main

import (
	"context"
	// "log"
	"os"
	"fmt"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

const (
	bucket = "bkt002"
	KAFKA_URL = "localhost:9092"
	KAFKA_TOPIC = "multicloud-input"
	KAFKA_GROUP_ID = "TestGroupID"
	// KAFKA_URL = "KAFKA_URL"
	// KAFKA_TOPIC = "KAFKA_TOPIC"
	// KAFKA_GROUP_ID = "KAFKA_GROUP_ID"
)

var inps = [12] string {
	"photo1.txt",
	"photo2.txt",
	"photo3.txt",
	"photo4.txt",
	"photo5.txt",
	"photo6.txt",
	"photo7.txt",
	"photo8.txt",
	"photo9.txt",
	"photo10.txt",
	"photo11.txt",
	"photo12.txt",

	// photo1.txt apple
	// photo2.txt zebra
	// photo3.txt parrot
	// photo4.txt orange
	// photo5.txt pear
	// photo6.txt lion
	// photo7.txt cat
	// photo8.txt dog
	// photo9.txt banana
	// photo10.txt duck
	// photo11.txt dove
	// photo12.txt mango
}



func getKafkaWriter(kafkaURL, topic, groupID string) *kafka.Writer {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}



func upload() {

	for _, fname := range inps  {
		fmt.Println("Uploading [", bucket, "] [", fname, "]")
		GelatoUpload(bucket, fname)
	}

}

func processFile(filename string) []byte {

	// fmt.Println("Reading input file:", filename)
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error Opening input file", filename, err)
		return nil
	}
	
	defer f.Close()

	buf := make([]byte, 16)
    n, err := f.Read(buf)
	if err != nil {
		fmt.Println("Error Reading input file", filename, err)
		return nil
	}

	if n > 0 {
		// fmt.Println("[", n, "]", string(buf))
	}

	return buf
}

func main() {
	// get kafka reader using environment variables.
	// kafkaURL, ok := os.LookupEnv(KAFKA_URL)
	// if !ok {
	// 	fmt.Println("Error getting KAFKA URL from environment variable")
	// }

	// topic, ok := os.LookupEnv(KAFKA_TOPIC)
	// if !ok {
	// 	fmt.Println("Error getting KAFKA TOPIC from environment variable")
	// }

	// groupID, ok := os.LookupEnv(KAFKA_GROUP_ID)
	// if !ok {
	// 	fmt.Println("Error getting KAFKA GROUP ID from environment variable")
	// }

	kafkaURL := KAFKA_URL
	topic := KAFKA_TOPIC
	groupID := KAFKA_GROUP_ID

	writer := getKafkaWriter(kafkaURL, topic, groupID)
	defer writer.Close()


	for _, fname := range inps  {
		fmt.Println("Downloading [" + bucket + "] [" + fname + "]")
		GelatoDownload(bucket, fname)

		value := string (processFile(fname))
		values := strings.Fields(value)
		value = values[0]

		fmt.Println("filename =", fname, "Value =", value)
		writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fname),
				Value: []byte(value),
			},
		)
	}
}
