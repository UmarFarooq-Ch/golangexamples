package golangexamples

import (
	"github.com/ehteshamz/greetings"
)

func ConcatSlice(sliceToConcat []byte) string {
	var concatString string
	for i := 0; i < len(sliceToConcat); i++ {
		concatString += string(sliceToConcat[i])
		if i != (len(sliceToConcat) - 1) {
			concatString += "-"
		}
	}
	return concatString
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	for i := 0; i < len(sliceToEncrypt); i++ {
		sliceToEncrypt[i] += byte(ceaserCount - 48)
	}
}

func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
