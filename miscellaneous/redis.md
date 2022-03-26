# Redis

```bash
docker container exec -it redis sh

docker container exec -it redis redis-cli KEYS \*

docker container exec -it redis redis-cli GET "/recommendations/17f9ad3783d07a4f3e3db8ce"

docker container exec -it redis redis-cli GET "/user/bearer/17f9ad3783d07a4f3e3db8ce"

docker container exec -it redis redis-cli

$ redis-cli

KEYS *

GET "/recommendations/17f9ad3783d07a4f3e3db8ce"

DEL "/recommendations/17f9ad3783d07a4f3e3db8ce"

FLUSHALL

FLUSHDB

QUIT
```
