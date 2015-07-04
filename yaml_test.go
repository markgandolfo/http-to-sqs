package main_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var data = `
mappings:
	project1:
		path: /api/v1/pathname1
		sqs: my_sqs_queue1
	project2:
		path: /api/v1/pathname2
		sqs: my_sqs_queue2
`

var expectedString = `
{Mappings:map[project1:{Path:/api/v1/pathname1 Sqs:my_sqs_queue1} project2:{Path:/ap
i/v1/pathname2 Sqs:my_sqs_queue2}]}`

func TestParser(t *testing.T) {
	assert.True(t, parseYaml(data), expectedString)
}
