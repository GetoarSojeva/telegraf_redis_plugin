
## How to install plugins:
```bash session
foo@bar:~$ git clone https://github.com/influxdata/telegraf.git
foo@bar:~$ git clone git@github.com:GetoarSojeva/telegraf_redis_plugin.git

foo@bar:~$ cd ./telegraf/plugins/outputs/
foo@bar:~$ mkdir redis_output

foo@bar:~$ cp -r ./telegraf_redis_plugins/plugins/outputs/redis_output/* ./telegraf/plugins/outputs/redis_output
foo@bar:~$ cp -r ./telegraf_redis_plugins/plugins/outputs/all/* ./telegraf/plugins/outputs/all/

foo@bar:~$ cd ./telegraf/plugins/inputs/
foo@bar:~$ mkdir redis_input

foo@bar:~$ cp ./telegraf_redis_plugins/plugins/inputs/redis_output/* ./telegraf/plugins/inputs/redis_input
foo@bar:~$ cp ./telegraf_redis_plugins/plugins/inputs/all/* ./telegraf/plugins/inputs/all/

foo@bar:~$ cd ./telegraf

foo@bar:~$ make
foo@bar:~$ ./telegraf
```



## Telegraf Config - telegraf.conf

```toml
[[inputs.redis_input]]
  servers = ["127.0.0.1:6379"]
  queue = "myqueue_input"


[[outputs.redis_output]]
  servers = ["127.0.0.1:6379"]
  queue = "myqueue_output"
```
