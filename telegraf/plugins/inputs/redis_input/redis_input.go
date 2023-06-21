package redis_input

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type RedisInput struct {
	Servers []string `toml:"servers"`
	Queue   string   `toml:"queue"`
}

func (r *RedisInput) Description() string {
	return "Reads values from a Redis queue using the RPOP method"
}

func (r *RedisInput) SampleConfig() string {
	return `
  ## Redis server(s)
  servers = ["tcp://localhost:6379"]
  ## Redis queue name
  queue = "queue_name"
`
}

func (r *RedisInput) Gather(acc telegraf.Accumulator) error {

	conn, err := redis.Dial("tcp", r.Servers[0])
	if err != nil {
		return fmt.Errorf("unable to connect to Redis server: %v", err)
	}
	defer conn.Close()

	for {
		reply, err := conn.Do("RPOP", r.Queue)
		if err != nil {
			return fmt.Errorf("Error reading from Redis queue %s: %s", r.Queue, err)
		}

		value, ok := reply.([]byte)
		if ok {
			// Send the value to Telegraf as a metric
			fields := map[string]interface{}{
				"value": string(value),
			}
			acc.AddFields("redis_queue", fields, nil)
		} else if reply == nil {
			// Queue is empty
			break
		} else {
			return fmt.Errorf("Unexpected reply type from Redis: %T", reply)
		}
	}
	return nil
}

func init() {
	inputs.Add("redis_input", func() telegraf.Input { return &RedisInput{} })
}
