package q04

import "strings"

//Assume the string only contains lowercase letters
func IsPalindromePermutation(s string) bool {

	counts, total := countCharacters(s)

	if total%2 == 0 {
		//If len is even, all character counts must be EVEN
		for _, count := range counts {
			if count%2 == 1 {
				return false
			}
		}
		return true

	} else {
		//If len is odd, exactly one character count must be ODD
		foundOddCount := false
		for _, count := range counts {
			if count%2 == 1 {
				if foundOddCount {
					return false
				}
				foundOddCount = true
			}
		}
		return foundOddCount
	}

}

func countCharacters(s string) ([]int, int) {
	total := 0
	counts := make([]int, 26) //only need to consider lowercase letters

	s = strings.ToLower(s)

	for i := 0; i < len(s); i++ {
		currentChar := s[i]
		if currentChar >= 'a' && currentChar <= 'z' {
			index := currentChar - 'a'
			counts[index]++
			total++
		}
	}
	return counts, total
}
