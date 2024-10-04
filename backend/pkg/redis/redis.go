package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedis(url string) (*redis.Client, error) {
	opts := &redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	}

	client := redis.NewClient(opts)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Обработчик, принимающий запрос от фронтенда и публикующий задачу в Redis
//func (h *APIHandler) StartTask(c echo.Context) error {
//	var task RequestBody
//	if err := c.Bind(&task); err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
//	}
//
//	// Преобразуем задачу в JSON для отправки в Redis
//	taskData, err := json.Marshal(task)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to serialize task"})
//	}
//
//	// Генерация ID задачи для её отслеживания
//	taskID := fmt.Sprintf("task_%d", time.Now().UnixNano())
//
//	// Публикация задачи в очередь Redis
//	err = h.redisClient.LPush(context.Background(), "task_queue", taskData).Err()
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to publish task"})
//	}
//
//	// Возвращаем taskID, чтобы фронтенд мог отслеживать статус
//	return c.JSON(http.StatusAccepted, map[string]string{"task_id": taskID})
//}
