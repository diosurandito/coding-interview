package slidingwindow

import "math"

func fruitIntoBasket(trees []int) int {
	treeTypesMap := make(map[int]bool)
	var max, start int

	for end, treeType := range trees {
		if len(treeTypesMap) < 2 && !treeTypesMap[treeType] {
			treeTypesMap[treeType] = true
			max = int(math.Max(float64(max), float64(end - start + 1)))
		} else if treeTypesMap[treeType] {
			max = int(math.Max(float64(max), float64(end - start + 1)))
		} else {
			treeTypesMap = make(map[int]bool)
			treeTypesMap[trees[end-1]] = true
			treeTypesMap[treeType] = true
			start = end - 1 

			for trees[start] == trees[start-1] {
				start--
			}

			max = int(math.Max(float64(max), float64(end-start+1)))
		}
	}
	return max
}