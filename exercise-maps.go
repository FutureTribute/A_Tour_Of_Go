// moretypes/23
package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	res := make(map[string]int)
	for _, v := range words {
		if _, ok := res[v]; ok == true {
			res[v] += 1
		} else {
			res[v] = 1
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
