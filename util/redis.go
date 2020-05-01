package util

import (
	"fmt"

	"github.com/go-redis/redis"
)

//RedisNewClient return cl to con
func RedisNewClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	//fmt.Println(pong, err)
	// Output: PONG <nil>
	return client, nil
}

//RedidTestClient simple test to check Connection Work or not
func RedidTestClient() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		return err
	}

	val, err := client.Get("key").Result()
	if err != nil {
		return err
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		return err
	} else {
		fmt.Println("key2", val2)
	}
	res := client.HSet("Mylist", "f1", "1")
	if res.Err() != nil {
		return res.Err()
	}
	valHget := client.HGet("Mylist", "f1")

	if valHget.Err() != nil {
		return valHget.Err()
	}
	fmt.Println("my list f1:", valHget.Val())
	return nil
}
