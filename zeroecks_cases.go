package main

import (
	"C"
	"testing"
	"github.com/stretchr/testify/assert"
)

func testVersion(t *testing.T) {
	var output *C.char = (*C.char)(C.malloc(1024 * C.sizeof_char))
	RVExtension(output, 1024, "version")
	assert.Equal(t, "1.0", C.GoString(output))
}

func testHit(t *testing.T) {

}

func testOperation(t *testing.T) {

}

func testPlayer(t *testing.T) {

}
