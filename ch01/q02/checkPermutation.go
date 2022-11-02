package q02

//Get character counts for each string, then compare those counts for equality
//time complexity = O(a + b)
//space complexiry = O(1) (character sets are fixed)
func CheckPermutation(a, b string) bool {

	aCounts := countCharacters(a)
	bCounts := countCharacters(b)

	for i := 0; i < 128; i++ {
		if aCounts[i] != bCounts[i] {
			return false
		}
	}
	return true
}

func countCharacters(s string) []int {
	counts := make([]int, 128) //Assumption: ASCII (128 characters)

	for i := 0; i < len(s); i++ {
		currentChar := s[i]
		counts[currentChar]++
	}
	return counts
}

//Count the characters in string a
//then decrement the counts as we iterate through string b
func CheckPermutation2(a, b string) bool {

	//check that lengths are equal
	if len(a) != len(b) {
		return false
	}

	counts := countCharacters(a)

	for i := 0; i < len(b); i++ {
		currentChar := b[i]
		counts[currentChar]--
		if counts[currentChar] < 0 {
			return false
		}
	}

	return true
}
