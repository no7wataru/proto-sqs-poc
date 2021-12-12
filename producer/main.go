package main

import (
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	pb "github.com/no7wataru/proto-sqs-poc/proto"
	"google.golang.org/protobuf/proto"
)

const (
	endpoint  = "http://localhost:9324"
	region    = "ap-northeast-1"
	queueName = "proto-test-mq"
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	client := sqs.New(sess, &aws.Config{
		Endpoint: aws.String(endpoint),
		Region:   aws.String(region),
	})

	output, err := client.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		panic(err)
	}
	println(*output.QueueUrl)

	m := &pb.Message{
		Title: "This is the title of the message",
		Sentences: []*pb.Sentence{
			{
				Id:   1,
				Text: "This is the first sentence.",
			},
			{
				Id:   2,
				Text: "This is the second sentence.",
			},
		},
	}
	data, _ := proto.Marshal(m)
	message := base64.StdEncoding.EncodeToString(data)

	res, err := client.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    output.QueueUrl,
		MessageBody: aws.String(message),
	})
	if err != nil {
		panic(err)
	}

	println(*res.MessageId)
}
