package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAllLowerCase(t *testing.T) {
	type testCase struct {
		name           string
		input          string
		expectedResult bool
		hasError       bool
	}

	testTable := []testCase{
		{name: "all lowercase string input",
			input:          "aabbcc",
			expectedResult: true,
			hasError:       false},
		{name: "some upper case input",
			input:          "aAbBcC",
			expectedResult: false,
			hasError:       false},
		{name: "contains a number",
			input:          "aabbcc1",
			expectedResult: false,
			hasError:       true},
	}

	for _, test := range testTable {
		actual, err := IsAllLowerCase(test.input)

		assert.Equal(t, test.expectedResult, actual, test.name)

		if test.hasError {
			assert.NotNil(t, err, test.name)
		} else {
			assert.Nil(t, err, test.name)
		}
	}
}
