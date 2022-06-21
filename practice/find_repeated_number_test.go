package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRepeatedNumber(t *testing.T) {
	expected := []interface{}{2,8}
	observed := findRepeatedNumber([]interface{}{3,4,2,8,1,2,8})
	assert.Equal(t, expected, observed, "Both should be equal!")
}
