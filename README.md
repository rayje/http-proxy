# http-proxy
## Overview
A configurable HTTP proxy that can listen for HTTP requests and respond with status/messages based on configuration.

## Installation
To install 

```
go get github.com/rayje/http-proxy
```

## Getting Started
The HTTP proxy will setup handlers to listen over a specific port for calls made to HTTP endpoints. Based on the configuration, the endpoint can respond with defined status codes, headers, and body.

### Configuration
The proxy uses a JSON based configuration to define the port the proxy should listen, as well as the endpoints.

The following JSON shows an example of a basic JSON configuration:
```JSON
{
    "port": 8080,
    "endpoints": [
        { 
            "name": "/test",
            "status": 200,
            "headers": [
                {"X-Test":"Test Header"}
            ],
            "body": "Test body",
        },
    ]
}
```

Based on the previous config, the HTTP proxy will setup an endpoint to listen on port 8080 and handle requests for the "/test" endpoint. Requests made to this endpoint will receive the follwing HTTP response:
```
HTTP/1.1 200 OK
X-Test: Test Header
Date: Sun, 02 Aug 2015 03:00:25 GMT
Content-Length: 9
Content-Type: text/plain; charset=utf-8

Test Body
```

### Running
```
./http-proxy --config my-config.json
```

