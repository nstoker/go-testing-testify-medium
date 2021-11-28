package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	inputA := 1
	inputB := 2

	expected := 3

	actual := Add(inputA, inputB)

	assert.Equal(t, expected, actual)
}
