package main

import (
	"fmt"
	"sort"
)

type ByLen []string

func (a ByLen) Len() int {
	return len(a)
}

func (a ByLen) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func (a ByLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {

	s := []string{"a", "bb", "ccc", "d", "ee", "fff"}

	fmt.Println(s)

	sort.Sort(ByLen(s))

	fmt.Println(s)
}
