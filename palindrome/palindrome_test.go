package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	expected := true
	observed1 := IsPalindrome("a")
	observed2 := IsPalindrome("aba")
	observed3 := IsPalindrome("katak")

	assert.Equal(t, expected, observed1, "Both should be equal!")
	assert.Equal(t, expected, observed2, "Both should be equal!")
	assert.Equal(t, expected, observed3, "Both should be equal!")
}

func TestIsNotPalindrome(t *testing.T) {
	observed1 := IsPalindrome("as")
	observed2 := IsPalindrome("abn")
	observed3 := IsPalindrome("kathak")

	assert.False(t, observed1, "Both should be equal!")
	assert.False(t, observed2, "Both should be equal!")
	assert.False(t, observed3, "Both should be equal!")
}

func TestIsPalindromeWithoutTemp(t *testing.T) {
	observed1 := IsPalindromeWithoutTemp("a")
	observed2 := IsPalindromeWithoutTemp("aba")
	observed3 := IsPalindromeWithoutTemp("katak")

	assert.True(t, observed1, "Both should be equal!")
	assert.True(t, observed2, "Both should be equal!")
	assert.True(t, observed3, "Both should be equal!")
}

func TestIsNotPalindromeWithoutTemp(t *testing.T) {
	observed1 := IsPalindromeWithoutTemp("as")
	observed2 := IsPalindromeWithoutTemp("abn")
	observed3 := IsPalindromeWithoutTemp("kathak")

	assert.False(t, observed1, "Both should be equal!")
	assert.False(t, observed2, "Both should be equal!")
	assert.False(t, observed3, "Both should be equal!")
}

func TestIsPalindromeWithRecursive(t *testing.T) {
	observed1 := IsPalindromeWithRecursive("a")
	observed2 := IsPalindromeWithRecursive("aba")
	observed3 := IsPalindromeWithRecursive("katak")

	assert.True(t, observed1, "Both should be equal!")
	assert.True(t, observed2, "Both should be equal!")
	assert.True(t, observed3, "Both should be equal!")
}

func TestIsNotPalindromeWithRecursive(t *testing.T) {
	observed1 := IsPalindromeWithRecursive("as")
	observed2 := IsPalindromeWithRecursive("abn")
	observed3 := IsPalindromeWithRecursive("kathak")

	assert.False(t, observed1, "Both should be equal!")
	assert.False(t, observed2, "Both should be equal!")
	assert.False(t, observed3, "Both should be equal!")
}
