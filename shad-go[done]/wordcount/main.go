//go:build !solution

package main

import (
	"fmt"
	"os"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	linesMap := make(map[string]int)
	argsWithoutProg := os.Args[1:]

	for _, s := range argsWithoutProg {
		data, err := os.ReadFile(s)
		checkError(err)
		arr := strings.Split(string(data), "\n")
		for _, v := range arr {
			linesMap[v]++
		}
	}

	for key, value := range linesMap {
		if value > 1 {
			fmt.Printf("%s\t%s\n", fmt.Sprint(value), key)
		}
	}
}
