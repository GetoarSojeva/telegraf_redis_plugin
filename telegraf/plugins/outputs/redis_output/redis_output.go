package redis_output

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/outputs"
)

type RedisOutput struct {
	Servers []string `toml:"servers"`
	Queue   string   `toml:"queue"`
}

func (r *RedisOutput) Connect() error {
	return nil
}

func (r *RedisOutput) Close() error {
	return nil
}

func (r *RedisOutput) SampleConfig() string {
	return `
  ## Redis server addresses
  servers = ["localhost:6379"]

  ## Redis queue name
  queue = "myqueue"
`
}

func (r *RedisOutput) Description() string {
	return "Send metrics to Redis using RPOP and RPUSH methods"
}

func (r *RedisOutput) Write(metrics []telegraf.Metric) error {
	conn, err := redis.Dial("tcp", r.Servers[0])
	if err != nil {
		return fmt.Errorf("unable to connect to Redis server: %v", err)
	}
	defer conn.Close()

	for _, metric := range metrics {
		value := metric.Fields()["value"].(string)
		// fmt.Println(value)

		_, err = conn.Do("RPUSH", r.Queue, value)
		if err != nil {
			return fmt.Errorf("error pushing metric to Redis: %v", err)
		}
	}

	return nil
}

func init() {
	outputs.Add("redis_output", func() telegraf.Output { return &RedisOutput{} })
}
