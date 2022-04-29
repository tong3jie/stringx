package stringx

func Includes(needle string, haystack []string) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}

func At(needle string, haystack []string) int {
	for i := range haystack {
		if haystack[i] == needle {
			return i
		}
	}
	return -1
}
