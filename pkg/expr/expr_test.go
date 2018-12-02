package expr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode_String(t *testing.T) {
	input := "15 7 1 1 + - / 3 * 2 1 1 + + -"

	result, err := ParseString(input)

	assert.NoError(t, err)
	assert.Equal(t, "(((15 / (7 - (1 + 1))) * 3) - (2 + (1 + 1)))", result.String())
}

func TestParseString(t *testing.T) {
	str := "3 4 5 + -"

	result, err := ParseString(str)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestNode_Evaluate(t *testing.T) {
	input := "15 7 1 1 + - / 3 * 2 1 1 + + -"

	node, err := ParseString(input)

	assert.NoError(t, err)

	result, err := node.Evaluate()

	assert.NoError(t, err)
	assert.Equal(t, 5.0, result)
}