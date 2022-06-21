package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSizeSubarraySum(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7

	expected := 2
	observed := minSizeSubarraySum(target, nums)
	assert.Equal(t, expected, observed, "Both should be equal!")
}