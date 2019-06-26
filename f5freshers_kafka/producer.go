package f5freshers_kafka
import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
//	"github.com/confluentinc/confluent-kafka-go/kafka"	

)



func Produce(messages []string,topic string) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

//	topic := "myTopic"
	for _, word := range messages {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
	//fmt.Println("published")
	p.Flush(15 * 1000)
}
