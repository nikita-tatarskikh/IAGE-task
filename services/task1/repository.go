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
	log.Println(request.GetKey(), request.GetValue())
	r.rdb.Incr(ctx, request.GetKey())
	val, err := r.rdb.Get(ctx, request.GetKey()).Int()
	log.Printf("val %d", val)
	if err != nil{
		log.Println(err)
	}
	incrementedValue.SetValue(val)
	log.Printf("incrementedValue %+v\n",incrementedValue)
	return &incrementedValue
}

func NewRepository(rdb *redis.Client) *RedisRepository {
	return &RedisRepository{rdb: rdb}
}
