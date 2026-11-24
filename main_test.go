package main

import (
	"reflect"
	"testing"
)

func TestZArray(t *testing.T) {
	cases := []struct {
		in   string
		want []int
	}{
		{"aabxaab", []int{7, 1, 0, 0, 3, 1, 0}},
		{"aaaa", []int{4, 3, 2, 1}},
		{"abcd", []int{4, 0, 0, 0}},
		{"", []int{}},
		{"a", []int{1}},
	}
	for _, c := range cases {
		if got := ZArray(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("ZArray(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestZSearch(t *testing.T) {
	cases := []struct {
		text, pattern string
		want          []int
	}{
		{"aabxaabxcaabxaabxay", "aabx", []int{0, 4, 9, 13}},
		{"aaaa", "aa", []int{0, 1, 2}},
		{"abc", "d", []int{}},
		{"abc", "", []int{}},
	}
	for _, c := range cases {
		if got := Search(c.text, c.pattern); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Search(%q,%q) = %v, want %v", c.text, c.pattern, got, c.want)
		}
	}
}
