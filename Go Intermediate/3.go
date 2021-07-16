package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'customSorting' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY strArr as parameter.
 */

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

func customSorting(strArr []string) []string {
	sort.Strings(strArr)
	sort.Sort(ByLen(strArr))
	return strArr
}

func main() {
	fmt.Println(customSorting([]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}))
	//[]string {"Earth", "Venus", "Jupiter", "Mercury", "Neptune", "Saturn", "Uranus", "Mars"}
}
