package q01

import "strings"

// Iterate through the characters in the string
// Use a hashmap to track if we've seen a character before
// time complexity = O(n)
// space complexity = O(n) (maybe O(1) since it won't exceed length of character set)
func IsUnique(s string) bool {

	visited := map[byte]bool{}

	for i := 0; i < len(s); i++ {
		currentChar := s[i]

		if _, exists := visited[currentChar]; exists {
			return false
		}
		visited[currentChar] = true
	}
	return true
}

//Don't use additional data structures
//Need to compare each letter to every other letter
// time complexity = O(n^2)
// space complexity = O(1)
func IsUnique2(s string) bool {

	if len(s) <= 1 {
		return true
	}

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

//Use Boolean Array instead of a HashMap
//Assumes there are only 128 possible characters (ASCII)
// time complexity = O(n)
// space complexity = O(1)
func IsUnique3(s string) bool {

	//If the string length > 128, then some character MUST repeat
	if len(s) > 128 {
		return false
	}

	//bool is initialized as false (aka NOT visited)
	visited := make([]bool, 128)

	for i := 0; i < len(s); i++ {
		currentChar := s[i]

		if visited[currentChar] {
			return false
		}
		visited[currentChar] = true
	}
	return true
}

// Use BitVector instead of a Boolean Array
// Assumes there are no more than 32 possible characters (we'll only look at lowercase alphacharacters)
// time complexity = O(n)
// space complexity = O(1)
func IsUnique4(s string) bool {

	var bitvector int32 //use int32 with 32 bits

	//Convert uppercase letters to lowercase
	s = strings.ToLower(s)

	for i := 0; i < len(s); i++ {
		currentChar := s[i]

		//ignore any characters that are not lowercase alpha
		if currentChar < 'a' || currentChar > 'z' {
			break
		}

		index := currentChar - 'a'
		mask := int32(1 << index)
		if (bitvector & mask) == mask {
			//the bit is ON, so we've seen this character already
			return false
		}
		//flip the bit to ON to mark that saw this character
		bitvector = bitvector | mask
	}
	return true
}
