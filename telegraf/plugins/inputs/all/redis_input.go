//go:build !custom || inputs || inputs.redis_input

package all

import _ "github.com/influxdata/telegraf/plugins/inputs/redis_input" // register plugin

