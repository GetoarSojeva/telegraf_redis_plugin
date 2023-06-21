
## How to install plugins:
```
git clone https://github.com/influxdata/telegraf.git

cd ./telegraf/plugins/outputs/
mkdir redis_output


cp -r telegraf_redis_plugins/plugins/outputs/redis_output/* ./telegraf/plugins/outputs/redis_output
cp -r telegraf_redis_plugins/plugins/outputs/all/* ./telegraf/plugins/outputs/all/



cd ./telegraf/plugins/inputs/
mkdir redis_input

cp telegraf_redis_plugins/plugins/inputs/redis_output/* ./telegraf/plugins/inputs/redis_input
cp telegraf_redis_plugins/plugins/inputs/all/* ./telegraf/plugins/inputs/all/


cd ./telegraf

make
```



## Telegraf Config

```toml
[[inputs.redis_input]]
  servers = ["127.0.0.1:6379"]
  queue = "myqueue_input"


[[outputs.redis_output]]
  servers = ["127.0.0.1:6379"]
  queue = "myqueue_output"
```
