package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := flag.String("s", "", "The string to check if it is a palindrome")
	flag.Parse()

	regex, _ := regexp.Compile("[^a-zA-Z0-9]+")

	processedString := strings.ToLower(regex.ReplaceAllString(*s, ""))
	reversedString := Reverse(processedString)

	if strings.Compare(processedString, reversedString) == 0 {
		fmt.Printf("👍 \"%s\" is a palindrome!\n", *s)
	} else {
		fmt.Printf("👎 \"%s\" is NOT a palindrome\n", *s)
	}
}

/**
 * Reverse string code from https://github.com/golang/example/blob/master/stringutil/reverse.go
 */
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
