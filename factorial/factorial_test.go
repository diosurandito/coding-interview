package factorial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	assert.Equal(t, 1, Factorial(0))
	assert.Equal(t, 1, Factorial(1))
	assert.Equal(t, 2, Factorial(2))
	assert.Equal(t, 6, Factorial(3))
	assert.Equal(t, 24, Factorial(4))
	assert.Equal(t, 120, Factorial(5))
}

func TestFactorialRecursive(t *testing.T) {
	assert.Equal(t, 1, FactorialRecursive(0))
	assert.Equal(t, 1, FactorialRecursive(1))
	assert.Equal(t, 2, FactorialRecursive(2))
	assert.Equal(t, 6, FactorialRecursive(3))
	assert.Equal(t, 24, FactorialRecursive(4))
	assert.Equal(t, 120, FactorialRecursive(5))
}

func TestFactorialTailRecursive(t *testing.T) {
	assert.Equal(t, 1, FactorialTailRecursive(1, 0))
	assert.Equal(t, 1, FactorialTailRecursive(1, 1))
	assert.Equal(t, 2, FactorialTailRecursive(1, 2))
	assert.Equal(t, 6, FactorialTailRecursive(1, 3))
	assert.Equal(t, 24, FactorialTailRecursive(1, 4))
	assert.Equal(t, 120, FactorialTailRecursive(1, 5))
}
