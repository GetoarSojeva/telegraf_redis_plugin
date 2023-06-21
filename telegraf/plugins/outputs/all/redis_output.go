//go:build !custom || outputs || outputs.redis

package all

import _ "github.com/influxdata/telegraf/plugins/outputs/redis_output" // register plugin

