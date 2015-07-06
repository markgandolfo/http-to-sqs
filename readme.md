### HTTP → SQS Queue 
[![Build](https://travis-ci.org/markgandolfo/http-to-sqs.svg?branch=master)](https://travis-ci.org/markgandolfo/http-to-sqs)

A simple golang service that will accept a http/https request at a specific url, and send the params to an sqs queue.

Configuration is in JSON

- Path is the HTTP request path
- Queue is the name of the SQS Queue

```JSON
[
    {
	"Path": "/api/v1/pathname1",
	"Queue": "my_sqs_queue1"
    },
    {
	"Path": "/api/v1/pathname2",
	"Queue": "my_sqs_queue2"
    }
]
```

Another config file will be checked for aws credentials, if it's not present, a check in the ENV for `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`

#### Suggested phase 2

* HTTP/HTTPS support
* HTTP Basic Auth support
