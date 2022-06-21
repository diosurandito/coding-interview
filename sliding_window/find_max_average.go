package slidingwindow

import "math"

//naif Solution
func findMaxAverage(nums []int, k int) float64 {
	max := math.Inf(-1)

	for i := 0; i < len(nums); i++ {
		naifSum := 0
		for j := 0; j < k; j++ {
			if !(i+(k-1) > len(nums)-1) {
				naifSum += nums[i+j]
			} else {
				i = len(nums)
				j = k
			}
		}
		if naifSum != 0 {
			max = math.Max(max, float64(naifSum)/float64(k))
		}
	}
	if max == math.Inf(-1) {
		return 0 
	}
	return max
}

//Good Solution with Sliding window
func findMaxAverageSlidingWindow(nums []int, k int) float64 {
	var windowSum int
	var start int
	max := math.Inf(-1)

	for end := 0; end < len(nums); end++ {
		windowSum += nums[end]

		if (end - start + 1) == k {
			// calculate average
			max = math.Max(max, float64(windowSum)/float64(k))
			windowSum -= nums[start]
			start += 1
		}
	}
	return max
} 

