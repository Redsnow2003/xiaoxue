package service

import (
	"fmt"
	"main/redis"

	redigoRedis "github.com/gomodule/redigo/redis"
)

type RedisSubService struct {
}

func (r *RedisSubService) StartRedisSubService() {
	go func() {
		redisConn := redis.RedisPool.Get()
		defer redisConn.Close()
		psc := redigoRedis.PubSubConn{Conn: redisConn}
		psc.Subscribe("Alarm_system_warning")
		psc.Subscribe("Alarm_system_error")
		psc.Subscribe("Alarm_system_fatal")
		defer psc.Unsubscribe()
		for {
			switch r := psc.Receive().(type) {
			case error:
				fmt.Println("redis sub receive error:", r)
			case redigoRedis.Message:
				{
					switch r.Channel {
					case "Alarm_system_warning":
						fmt.Printf("redis sub receive message: %s %s\n", r.Channel, string(r.Data))
					case "Alarm_system_error":
						fmt.Printf("redis sub receive message: %s %s\n", r.Channel, string(r.Data))
					case "Alarm_system_fatal":
						fmt.Printf("redis sub receive message: %s %s\n", r.Channel, string(r.Data))
					default:
						fmt.Printf("redis sub receive message: %s %s\n", r.Channel, string(r.Data))
					}
				}
			}
		}
	}()
}
