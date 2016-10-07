package main

/*
#cgo CFLAGS: -std=c11
#include <stdlib.h>
#include <string.h>
 */
import "C"

import (
	"unsafe"
	"net/http"
	"github.com/dghubble/sling"
	"strings"
)

const version = "1.0"
const endpoint = "http://localhost:8080/api/"

//export RVExtension
func RVExtension(output *C.char, size int, function string) {
	splits := strings.SplitN(function, " ", 2)

	command := splits[0]

	switch command {
	case "version":
		buffer := C.CString(version)
		defer C.free(unsafe.Pointer(buffer))

		C.memcpy(unsafe.Pointer(output), unsafe.Pointer(buffer), (C.size_t)(size * C.sizeof_char))
	default:
		return
	}
}

func main() { }

type Client struct {
	template *sling.Sling
	Hits *HitService
	Operations *OperationService
	Players *PlayerService
}

func NewClient() *Client {
	agent := &http.Client {
		Transport: &TransportAdapter{},
	}

	template := sling.New().Client(agent).Base(endpoint).Set("Accept", "application/vnd.api+json")

	return &Client {
		template: template,
		Hits: newHitService(template.New()),
		Operations: newOperationService(template.New()),
		Players: newPlayerService(template.New()),
	}
}

type TransportAdapter struct {
	http.Transport
}

func (t *TransportAdapter) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := t.Transport.RoundTrip(req)
	if res != nil {
		res.Header.Set("Content-Type", "application/json")
	}
	return res, err
}
