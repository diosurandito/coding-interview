package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubStringWithoutRepeat(t *testing.T) {
	s := "pwwkew"
	expected := 3
	observed := longestSubStringWithoutRepeat(s)
	assert.Equal(t, expected, observed, "Both should be equal!")
}