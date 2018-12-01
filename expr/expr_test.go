package expr

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/rfaulhaber/rexpr/expr"
)

func TestNode_String(t *testing.T) {
	input := "15 7 1 1 + - / 3 * 2 1 1 + + -"

	result, err := expr.ParseString(input)

	assert.NoError(t, err)
	assert.Equal(t, "((15 / (7 − (1 + 1))) * 3) − (2 + (1 + 1))", result.String())
}

func TestParseString(t *testing.T) {
	str := "3 4 5 + -"

	result, err := expr.ParseString(str)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}