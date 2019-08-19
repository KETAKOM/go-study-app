package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var svc *sqs.SQS

const (
	AWS_REGION = os.Getenv("AWS_REGION")
	QUEUE_URL  = os.Getenv("QUEUE_URL")
)

func RetrieveMessage() error {
	params := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(QUEUE_URL),
		// 一度に取得する最大メッセージ数。最大でも10まで。
		MaxNumberOfMessages: aws.Int64(2),
		// これでキューが空の場合はロングポーリング(20秒間繋ぎっぱなし)になる。
		WaitTimeSeconds: aws.Int64(20),
	}
	resp, err := svc.ReceiveMessage(params)

	if err != nil {
		return err
	}

	fmt.Printf("messages count: %d\n", len(resp.Messages))

	if len(resp.Messages) == 0 {
		fmt.Println("empty queue.")
		return nil
	}

	// メッセージの数だけgoroutineを実行してみる。
	var wg sync.WaitGroup
	for _, m := range resp.Messages {
		wg.Add(1)
		go func(msg *sqs.Message) {
			defer wg.Done()
			fmt.Println(*msg.Body)
			if err := DeleteMessage(msg); err != nil {
				fmt.Println(err)
			}
		}(m)
	}

	wg.Wait()

	return nil
}

// メッセージを削除する。
func DeleteMessage(msg *sqs.Message) error {
	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(QUEUE_URL),
		ReceiptHandle: aws.String(*msg.ReceiptHandle),
	}
	_, err := svc.DeleteMessage(params)

	if err != nil {
		return err
	}
	return nil
}

func main() {
	sess := session.Must(session.NewSession())
	svc = sqs.New(sess, &aws.Config{Region: aws.String(AWS_REGION)})

	// 止めるまでメッセージを回収しつづける。
	for {
		if err := RetrieveMessage(); err != nil {
			log.Fatal(err)
		}
	}
}
