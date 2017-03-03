package main

import "strings"

func pop(slice []string) (string, []string) {
	ans := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return ans, slice
}

func isContain(s string, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
