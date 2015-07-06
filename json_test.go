package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testData = `
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
`

func TestParser(t *testing.T) {
	config := parseJSON(testData)
	assert.Equal(t, "/api/v1/pathname1", config[0].Path)
	assert.Equal(t, "my_sqs_queue1", config[0].Queue)

	assert.Equal(t, "/api/v1/pathname2", config[1].Path)
	assert.Equal(t, "my_sqs_queue2", config[1].Queue)
}
