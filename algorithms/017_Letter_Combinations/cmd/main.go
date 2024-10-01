package main

var result = []string{}

func letterCombinations(digits string) []string {

	result = []string{}

	letters := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	backtest(digits, "", 0, letters)

	return result
}

func backtest(digits, currentString string, pos int, letters map[string]string) {

	if digits == "" {
		return
	}
	if len(currentString) == len(digits) {
		result = append(result, currentString)
		return
	}

	auxLetter := letters[string(digits[pos])]

	pos += 1
	for _, letter := range auxLetter {
		backtest(digits, currentString+string(letter), pos, letters)
	}
}
