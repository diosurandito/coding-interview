package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadFindMaxAverage(t *testing.T) {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	expected := 12.75
	observed := findMaxAverage(nums, k)
	assert.Equal(t, expected, observed, "Both should be equal!")
}

func TestGoodFindMaxAverage(t *testing.T) {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	expected := 12.75
	observed := findMaxAverageSlidingWindow(nums, k)
	assert.Equal(t, expected, observed, "Both should be equal!")
}