package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	strarray := strings.Fields(s)
	var v int
	var ok bool
	for i := 0; i<len(strarray); i++{
		v,ok = m[strarray[i]]
		if ok == false{
			m[strarray[i]] = 1
		}else{
			m[strarray[i]] = v+1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
