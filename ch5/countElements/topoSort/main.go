package main

import (
	"fmt"
)

func main(){
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[sring][]string){
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string){
		for item := range m {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
			}
		}
	}

	visitAll()
}