# spacetraders-proxy

Easily throttle HTTP requests to the SpaceTrader api


## Usage
Make sure that `docker` & `docker-compose` are both installed on your machine

Then run the following (setting the SPACETRADES_TOKEN is not required since your clients may provide it however providing the SPACETRADERS_TOKEN means all requests going through the proxy will be authenticated)
```
export SPACETRADERS_TOKEN="<your-token>"
docker-compose up -d
```

At this point the service should be running in docker, to test it out make a curl statement
```
curl localhost:8081/v2/my/agent
```

Now you should see the response (something like)
```json
{
  "data": {
    "accountId": "blhnt...z0nl",
    "symbol": "COPE-TRADES",
    "headquarters": "X1-ZA40-15970B",
    "credits": 1110387
  }
}
```

Now wherever you implemented your client in your application, change the url from `https://api.spacetraders.io"` to `http://localhost:8081`. If your application is also containerized you may need to wire the networks together in your docker-compose and reference the service name instead of localhost.