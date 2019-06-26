package main

import(
	"fmt"
	kafka "./f5freshers_kafka"
)


func main(){


	messages := []string{"Message 1", "second message", "Message 3", "Fourth message"}
	
	topic:="kafka"
	

	kafka.Produce(messages,topic)

	fmt.Println("produced")

	var flag string

	scanning:fmt.Println("Press y to consume data, e to Exit")
	fmt.Scan(&flag)

	if flag=="y"{
		kafka.Consume(topic)
	}else if flag=="e"{
	}else{
		fmt.Println("Wrong Input, Try Again")
		goto scanning
	}
		
}
