package palindrome

import (
	"fmt"
	"strings"
)

func IsPalindrome(word string) bool {
	var temp string
	for i := len(word) - 1; i >= 0; i-- {
		temp = temp + string(word[i])
	}
	fmt.Println(temp)
	return strings.EqualFold(word, temp)
}

func IsPalindromeWithoutTemp(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		indexAwal := i
		indexAkhir := len(word) - i - 1
		compare := strings.EqualFold(string(word[indexAwal]), string(word[indexAkhir]))
		if !compare {
			return false
		}
	}
	return true
}

func isPalindromeRecursive(word string, i int) bool {
	if i < len(word)/2 {
		indexAwal := i
		indexAkhir := len(word) - i - 1
		compare := strings.EqualFold(string(word[indexAwal]), string(word[indexAkhir]))
		if !compare {
			return false
		} else {
			return isPalindromeRecursive(word, i+1)
		}
	} else {
		return true
	}
}

func IsPalindromeWithRecursive(word string) bool {
	return isPalindromeRecursive(word, 0)
}
