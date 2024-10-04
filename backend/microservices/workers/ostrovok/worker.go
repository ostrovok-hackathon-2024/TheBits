package ostrovok

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"

	queuePkg "github.com/ostrovok-hackathon-2024/The-Bits/backend/pkg/queue"
	redisPkg "github.com/ostrovok-hackathon-2024/The-Bits/backend/pkg/redis"
)

// INFO
// topic 1: receive task params (RequestBody) 	| name: worker.tasks
// topic 2: send task results 	(HotelInfo)		| name: worker.results

type worker struct {
	ctx                context.Context
	config             *Config
	httpClient         *http.Client
	queueWorkerResults queuePkg.Queue
	redisClient        *redis.Client
	consumerCh         chan []byte
	apiCallFrequency   time.Duration
}

type Worker interface {
	Run(context.Context) error
}

// TODO return error
func NewWorker(cfg *Config) Worker {
	redisClient, err := redisPkg.NewRedis(cfg.RedisURL)
	if err != nil {

	}

	ctx := context.Background()

	// TODO: to config
	queue := queuePkg.New(redisClient, "worker.tasks")
	consumerCh := make(chan []byte, 100)
	err = queue.Consume(ctx, consumerCh)
	if err != nil {
		// fatal
	}

	// TODO: to config
	queueWorkerResults := queuePkg.New(redisClient, "worker.results")

	return &worker{
		ctx:                ctx,
		httpClient:         http.DefaultClient,
		config:             cfg,
		redisClient:        redisClient,
		consumerCh:         consumerCh,
		queueWorkerResults: queueWorkerResults,
		apiCallFrequency:   time.Second * 1,
	}
}

func (w *worker) Run(ctx context.Context) error {

	for {
		select {
		case event := <-w.consumerCh:
			var requestBody RequestBody
			if err := json.Unmarshal(event, &requestBody); err != nil {
				log.Printf("Failed to unmarshal event: %v", err)
				continue
			}

			// TODO
			// Логика обработки задачи (например, запрос к внешнему API)
			//result := w.processTask(requestBody)

			//w.queueWorkerResults.Produce(ctx, result)
		case <-ctx.Done():
			fmt.Println("stop running by ctx.Done()")
			return nil
		}
	}
}

func (w *worker) start(ctx context.Context) error {
	ticker := time.NewTicker(w.apiCallFrequency)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")

			err := w.process()
			if err != nil {
				log.Printf("Worker: process error: %s", err.Error())
			}

		case <-ctx.Done():
			log.Printf("Worker: stop")
			return nil
		}
	}
}

func (w *worker) process() error {
	// 2: save results in queue
	//data, err := w.makeRequest()
	//if err != nil {
	//	return err
	//}

	//fmt.Println(len(data.Data.Hotels), data.Data.Hotels[0])
	//
	//rawHotelData := data.Data.Hotels[0]
	//
	//str, err := json.Marshal(rawHotelData)
	//
	//fmt.Println("\n")
	//fmt.Println(string(str), err)
	log.Fatalln("1")

	w.produce()

	return nil
}

func (w *worker) produce() {}

func (w *worker) makeRequest(_ *RequestBody) (*APIResponse, error) {
	// TODO: pass in func args
	requestBody := RequestBody{
		//CheckIn:     checkIn,
		//CheckOut:    checkOut,
		//Residency:   residency,
		//Language:    language,
		//Guests:      guests,
		//RegionId:    regionId,
		//Currency:    currency,
		//HotelsLimit: hotelsLimit,
	}

	requestBodyMarshaled, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	requestBodyRaw := bytes.NewReader(requestBodyMarshaled)

	req, err := http.NewRequest(http.MethodPost, w.config.APIURL, requestBodyRaw)
	if err != nil {
		log.Println("Failed to create request:", err)
		return nil, err
	}

	req.SetBasicAuth(w.config.Auth.Username, w.config.Auth.Password)
	req.Header.Add("Content-Type", "application/json")

	resp, err := w.httpClient.Do(req)
	if err != nil {
		log.Println("Request failed:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response body:", err)
		return nil, err
	}

	//TODO: сохранение результатов запроса

	// resultID := fmt.Sprintf("result_%d", time.Now().UnixNano())
	// err = w.redisClient.Set(context.Background(), resultID, body, 0).Err()
	// if err != nil {
	// 	return err
	// }

	responseBody := &APIResponse{}

	err = json.Unmarshal(body, responseBody)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

// Проверка завершенности задачи

// // APIHandler — обработчик для проверки статуса задачи
// func (h *APIHandler) GetTaskStatus(c echo.Context) error {
// 	taskID := c.Param("task_id")

// 	// Проверяем результат в Redis
// 	result, err := h.redisClient.Get(context.Background(), taskID).Result()
// 	if err == redis.Nil {
// 		return c.JSON(http.StatusNotFound, map[string]string{"status": "Task not found"})
// 	} else if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check task status"})
// 	}

// 	return c.JSON(http.StatusOK, map[string]string{"status": "Completed", "result": result})
// }
