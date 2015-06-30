### HTTP(s) -> SQS Queue 
![Build](https://travis-ci.org/markgandolfo/http-to-sqs.svg?branch=master)

A simple golang service that will accept a http request at a specific url, and send the params to an sqs queue.

There will be configuration via a config file (perhaps yaml), for example:

    ---
    project_name:
      path: /api/v1/create_article
      sqs: create_article

Another config file will be checked for aws credentials, if it's not present, a check in the ENV for `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`

#### Suggested phase 2

* HTTP/HTTPS support
* HTTP Basic Auth support
