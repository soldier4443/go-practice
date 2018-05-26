package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)
    strs := strings.Split(s, " ")
    
    for _, str := range strs {
        if _, exist := m[str]; exist {
            m[str] += 1
        } else {
        	m[str] = 1
        }
    }
    
    return m
}

func main() {
    wc.Test(WordCount)
}
