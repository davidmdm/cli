package util

// FilterStrings returns a new slice of strings by applying a predicate funtion to the original slice
func FilterStrings(sl []string, pred func(string, int) bool) []string {
	result := []string{}
	for i, e := range sl {
		if pred(e, i) {
			result = append(result, e)
		}
	}
	return result
}

// GetIndexes returns a slice containing the indexes of the elements that match the predicate function
func GetIndexes(sl []string, pred func(string) bool) []int {
	result := []int{}
	for i := range sl {
		if pred(sl[i]) {
			result = append(result, i)
		}
	}
	return result
}

// StripDashPrefix returns a new string with the dashes stripped from the prefix
func StripDashPrefix(str string) string {
	idx := -1
	for _, b := range []byte(str) {
		if b == '-' {
			idx++
		} else {
			break
		}
	}
	return str[idx+1:]
}
