package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"syscall"
	"testing"
)

func TestReadFile(t *testing.T) {
	tempFileContents := "The file"
	f, err := ioutil.TempFile("", "testmainconf")
	if err != nil {
		panic(err)
	}

	defer syscall.Unlink(f.Name())
	ioutil.WriteFile(f.Name(), []byte(tempFileContents), 0644)
	contents := readFile(f.Name())

	assert.Equal(t, contents, "The file")
}
