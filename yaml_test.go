package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testData = `
mappings:
  project1:
    path: /api/v1/pathname1
    sqs: my_sqs_queue1
  project2:
    path: /api/v1/pathname2
    sqs: my_sqs_queue2
`

var expectedString = `main.Config{Mappings:map[string]main.Route{"project1":main.Route{Path:"/api/v1/pathname1", Sqs:"my_sqs_queue1"}, "project2":main.Route{Path:"/api/v1/pathname2", Sqs:"my_sqs_queue2"}}}`

func TestParser(t *testing.T) {
	assert.Equal(t, parseYaml(testData), expectedString)
}
