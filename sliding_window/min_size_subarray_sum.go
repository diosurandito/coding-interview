package slidingwindow

import "math"

func minSizeSubarraySum(target int, nums []int) int {
	var sum, start int
	minLen := 0

	for end := 0; end < len(nums); end++ {
		sum += nums[end]

		for sum >= target {
			if minLen == 0 {
				minLen = end - start + 1
			}
			minLen = int(math.Min(float64(minLen), float64(end-start+1)))
			sum -= nums[start]
			start++
		}
	}
	if minLen == 0 {
		return 0
	}

	return int(minLen)

}