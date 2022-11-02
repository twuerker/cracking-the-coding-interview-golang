package q03

import "regexp"

func URLify(s string, trueLength int) string {

	countSpaces := 0
	for i := 0; i < trueLength; i++ {
		if s[i] == ' ' {
			countSpaces++
		}
	}

	newLength := trueLength + (2 * countSpaces)
	newString := make([]byte, newLength)
	newIndex := len(newString) - 1

	for i := trueLength - 1; i >= 0; i-- {

		if s[i] == ' ' {
			newString[newIndex] = '0'
			newString[newIndex-1] = '2'
			newString[newIndex-2] = '%'
			newIndex = newIndex - 3
		} else {
			newString[newIndex] = s[i]
			newIndex = newIndex - 1
		}

	}
	return string(newString)
}

//Use regex to replace ' ' with "%20"
func URLify2(s string, trueLength int) string {
	s = s[0:trueLength]
	r := regexp.MustCompile(" ")
	return r.ReplaceAllString(s, "%20")
}
