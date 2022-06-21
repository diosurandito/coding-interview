package practice

func ArrayCountValues(s []interface{}) map[interface{}]uint {
	r := make(map[interface{}]uint)
	for _, v := range s {
		if c, ok := r[v]; ok {
			r[v] = c + 1
		} else {
			r[v] = 1
		}
	}

	return r
}

func findRepeatedNumber(numbers []interface{}) []interface{} {
	var sameNumbers []interface{}
	checked_numbers := ArrayCountValues(numbers)
	for number, v := range checked_numbers {
		if v > 1 {
			sameNumbers = append(sameNumbers, number)
		}
	}
	return sameNumbers

}