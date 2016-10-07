package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient()

	assert.NotNil(t, client)
	assert.NotNil(t, client.Hits)
	assert.NotNil(t, client.Players)
	assert.NotNil(t, client.Operations)
}

func TestHit(t *testing.T) { testHit(t) }
func TestOperation(t *testing.T) { test}

