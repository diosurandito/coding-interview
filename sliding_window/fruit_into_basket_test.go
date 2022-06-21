package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFruitIntoBasket(t *testing.T) {
	trees := []int{0, 1, 6, 6, 4, 4, 6}
	expected := 5
	observed := fruitIntoBasket(trees)
	assert.Equal(t, expected, observed, "Both should be equal!")
}