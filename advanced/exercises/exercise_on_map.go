package main

import (
	"strings"
	"golang.org/x/tour/wc"
)
// count words freq in a sentence
func WordCount(s string) map[string]int {
	res := make(map[string]int)
	res2 := make(map[string]int)

	// Method 1 : using split function
	wordsSlice := strings.Split(s, " ")
	for _, v := range wordsSlice {
		res[v]++
	}

	/// Method 2 : iterating through each char
	var eachWord string
	for i, v := range s {

		if v != ' ' {
			eachWord = strings.Join([]string{eachWord, string(v)}, "")
			// eachWord += string(v) // another way to concatenate strings
			if i+1 == len(s) {
				res2[eachWord]++
			}
		} else {
			res2[eachWord]++
			eachWord = ""
		}
	}
	return res2
}

func main() {
	wc.Test(WordCount)
}
