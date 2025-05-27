package database

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/tylerjusfly/foodchatai/internal/models"
)



func RedisClient() *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

	return rdb
}

func PersistMessage(chatID, sender, msg string) error {
	var rdb = RedisClient()

    ctx := context.Background()

    message := models.ChatMessage{
        Sender:    sender,
        Message:   msg,
    }

    data, _ := json.Marshal(message)

	// fmt.Println(json.Unmarshal(data, ""))

    return rdb.RPush(ctx, chatID, data).Err()
}

func LoadChat(chatID string) ([]models.ChatMessage, error) {
	var rdb = RedisClient()
    ctx := context.Background()
    values, err := rdb.LRange(ctx, chatID, 0, -1).Result()
    if err != nil {
        return nil, err
    }

    var messages []models.ChatMessage
    for _, val := range values {
        var msg models.ChatMessage

        if err := json.Unmarshal([]byte(val), &msg); err == nil {
            messages = append(messages, msg)
        }
    }
    return messages, nil
}
