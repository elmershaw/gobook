package main

import (
	"fmt"
	"log"
)

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	visited := map[string]bool{}
	var visit func(item string)
	visit = func(item string) {
		if visited[item] {
			log.Fatal("loop existed!")
		}
		if !seen[item] {
			seen[item] = true
			visited[item] = true
			for _, prereq := range m[item] {
				visit(prereq)
			}
			order = append(order, item)
		}
	}
	for s := range m {
		visit(s)
		visited = map[string]bool{}
	}
	return order
}

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"linear algebra":        {"calculus"},
}
