# canary-news-bot

### Requirements
- redis

```bash
# docker run example

docker run --name redis \
    -v /tmp/data:/data \
    --network canary-news-net \
    --rm  -d redis
```

### Running
```bash
docker run --rm --name canary-news-bot \
    --network canary-news-net \
    -e REDIS_HOST=redis \
    -e REDIS_PORT=6379 \
    -e API_KEY=[BOT-API-KEY] \
    -e CHANNEL_ID=[CHANNEL-ID] \
    -e DEBOUNCE=2000 \
    iknpx/canary-news-bot
```

