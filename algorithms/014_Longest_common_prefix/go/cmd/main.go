package main

func longestCommonPrefix(s []string) string {

	if len(s) == 1 {
		return s[0]
	}

	w := getMinLen(s)
	prefixLen := 1
	res := []byte{}

	for prefixLen <= len(w) {
		val := isLetterInSlice(w, prefixLen, s)
		if !val {
			break
		}
		res = []byte(w[:prefixLen])

		prefixLen++
	}
	return string(res)
}

func isLetterInSlice(letter string, prefix int, s []string) bool {

	for i := range s {
		if letter[:prefix] != s[i][:prefix] {
			return false
		}
	}

	return true
}

func getMinLen(s []string) string {
	min := s[0]

	for i := 1; i < len(s); i++ {
		if len(s[i]) < len(min) {
			min = s[i]
		}
	}

	return min
}
