package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Obj struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func LearnRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var obj = new(Obj)
	obj = &Obj{
		ID:       1,
		Username: "John Doe",
		Email:    "johndoe@example.com",
	}
	marshaledEntity, _ := json.Marshal(&obj)
	err := rdb.Set(ctx, strconv.Itoa(int(obj.ID)), string(marshaledEntity), time.Hour).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, strconv.Itoa(int(obj.ID))).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(strconv.Itoa(int(obj.ID)), val)

	rdb.Del(ctx, strconv.Itoa(int(obj.ID)))
	val2, err := rdb.Get(ctx, strconv.Itoa(int(obj.ID))).Result()
	if err == redis.Nil {
		fmt.Println("deleted", val2)
	} 
	fmt.Println(rdb.Exists(ctx, strconv.Itoa(int(obj.ID))).Result())
	
	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
}
