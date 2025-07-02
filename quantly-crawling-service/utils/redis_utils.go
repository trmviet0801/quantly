package utils

import (
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func UnmarshallRedisReturn[T any](r *redis.FTSearchResult) ([]*T, error) {
	if r.Total == 0 {
		return nil, fmt.Errorf("zero matched in redis db")
	}

	var result []*T
	for _, doc := range r.Docs {
		item, err := MapToStruct[T](doc.Fields)
		if err != nil {
			return result, err
		}
		result = append(result, &item)
	}
	return result, nil
}

// Convert data queried from Redis (in String) to object
func UnmarshallRedisReturnString[T any](value string) ([]*T, error) {
	var result []T
	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		fmt.Println("here")
		OnError(err)
		return nil, err
	}
	return ToPointerArray(result), nil
}
