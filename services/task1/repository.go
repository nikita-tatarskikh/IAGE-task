package task1

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

type RedisRepository struct {
	rdb *redis.Client
}

func (r *RedisRepository) IncrementByKeyAndReturnValue(request *Request) *Response {
	var incrementedValue Response
	ctx := context.Background()
	r.rdb.Set(ctx, request.GetKey(), request.GetValue(), 0)
	r.rdb.Incr(ctx, request.GetKey())
	val, err := r.rdb.Get(ctx, request.GetKey()).Int()
	if err != nil{
		log.Println("Ошибка при выполнении операции в Redis", err)
		return nil
	}
	incrementedValue.SetValue(val)
	return &incrementedValue
}

func NewRepository(rdb *redis.Client) *RedisRepository {
	return &RedisRepository{rdb: rdb}
}
