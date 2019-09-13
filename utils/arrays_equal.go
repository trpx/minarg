package utils

func StringArraysEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for idx, el := range a {
		if el != b[idx] {
			return false
		}
	}
	return true
}
