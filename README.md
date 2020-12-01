[![Build Status](https://cloud.drone.io/api/badges/isalikov/canary.icu-bot/status.svg)](https://cloud.drone.io/isalikov/canary.icu-bot)

# canary.icu-bot

Canarian news bot

### Requirements
- redis

```bash
# docker run example

docker run --name redis \
    -v /tmp/data:/data \
    --network canary.icu-net \
    --rm  -d redis
```

### Running
```bash
docker run --rm --name canary.icu-bot \
    --network canary.icu-net \
    -e REDIS_HOST=redis \
    -e REDIS_PORT=6379 \
    -e API_KEY=[BOT-API-KEY] \
    -e CHANNEL_ID=[CHANNEL-ID] \
    -e DEBOUNCE=3000 \
    iknpx/canary.icu-bot
```

### Used Feeds
- [iestafeta.com](https://iestafeta.com)
- [russkoe-105fm.ru](http://russkoe-105fm.ru)
- [espanarusa.com](https://espanarusa.com)
