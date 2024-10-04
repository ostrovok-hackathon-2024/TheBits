package queue

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Queue interface {
	Produce(ctx context.Context, value []byte) (err error)
	Consume(ctx context.Context, ch chan<- []byte) (err error)
}

type queue struct {
	redisClient  *redis.Client
	topic        string
	readInterval time.Duration
}

var _ Queue = new(queue)

func New(redisClient *redis.Client, topic string) Queue {
	return &queue{
		redisClient:  redisClient,
		topic:        topic,
		readInterval: time.Millisecond * 150,
	}
}

func (q *queue) Consume(ctx context.Context, ch chan<- []byte) error {
	ticker := time.NewTicker(q.readInterval)

	go func() {
		for {
			select {
			case <-ticker.C:
			RPopLoop:
				for {
					value, err := q.redisClient.RPop(ctx, q.topic).Result()
					if err != nil {
						if errors.Is(err, redis.Nil) {
							break RPopLoop
						}

						log.Printf("error while RPop: %s", err.Error())
						continue
					}
					if len(value) == 0 {
						break RPopLoop
					}

					ch <- []byte(value)
				}
			case <-ctx.Done():
				fmt.Println("Stop consuming by ctx.Done()")
				return
			}
		}
	}()

	return nil
}

func (q *queue) Produce(ctx context.Context, value []byte) (err error) {
	err = q.redisClient.LPush(ctx, q.topic, string(value)).Err()
	if err != nil {
		return err
	}

	return
}

//func (o *queue) ProcessQueue() {
//for {
//	task, err := o.redisClient.RPop(context.Background(), "task_queue").Result()
//	if err == redis.Nil {
//		log.Println("No tasks in the queue")
//		time.Sleep(1 * time.Second)
//		continue
//	} else if err != nil {
//		log.Printf("Error getting task from queue: %v", err)
//		continue
//	}
//
//	// Десериализация задачи
//	var requestBody worker.RequestBody
//	err = json.Unmarshal([]byte(task), &requestBody)
//	if err != nil {
//		log.Printf("Failed to unmarshal task: %v", err)
//		continue
//	}
//
//	// Отправка задачи воркеру
//	//go o.dispatchToWorker(requestBody)
//}
//}

//func (o *queue) dispatchToWorker(task worker.RequestBody) {
//	log.Printf("Dispatching task for region %d", task.RegionId)
//	// Прописать логику вызова
//}
