package utils

import "strings"

//MakeExcited takes a phrase which is a string and transforms it to uppercase with an exclamation mark
func MakeExcited(phrase string) string {
	return strings.ToUpper(phrase) + "!"
}
