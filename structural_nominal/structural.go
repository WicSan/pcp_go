package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type SortByAge []Person

func (s SortByAge) Len() int {
	return len(s)
}
func (s SortByAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortByAge) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	sort.Sort(SortByAge(people))
	fmt.Println(people)
}
