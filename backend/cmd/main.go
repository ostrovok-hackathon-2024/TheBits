package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ostrovok-hackathon-2024/The-Bits/backend/microservices/api"
	"github.com/ostrovok-hackathon-2024/The-Bits/backend/microservices/processing"
	"github.com/ostrovok-hackathon-2024/The-Bits/backend/microservices/workers/ostrovok"
)

func main() {
	fmt.Println("Hello, world!")

	ctx := context.Background()
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		redisURL, ok := os.LookupEnv("REDIS_URL")
		if !ok {
			redisURL = "redis:6379"
		}
		config := &ostrovok.Config{
			APIKey:   "key",
			RedisURL: redisURL,
		}

		// TODO: update NewWorker
		//worker, err := ostrovok.NewWorker(config).Run(ctx)
		//worker.Run()

		err := ostrovok.NewWorker(config).Run(ctx)
		if err != nil {
			log.Fatalf("Worker start failed: %s", err.Error())
			return
		}
	}()

	go func() {
		ctx := context.Background()
		config := &processing.Config{}

		err := processing.NewProcessing(config).Run(ctx)
		if err != nil {
			log.Fatalf("Processing start failed: %s", err.Error())
			return
		}
	}()

	go func() {
		ctx2 := context.Background()
		config := &api.Config{}

		err := api.NewServerAPI(config).Run(ctx2)
		if err != nil {
			log.Fatalf("Server API start failed: %s", err.Error())
		}
	}()

	wg.Wait()
	log.Println("App shutdown")
}
