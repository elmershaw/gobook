package main

import "strings"

func expand(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "foo", f("foo"))
}
