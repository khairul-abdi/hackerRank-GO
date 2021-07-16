package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"WAlice": 2, "Cecil": 1, "Bob": 3}
	// strArr := []string{"WAlice", "Cecil", "Bob"}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	fmt.Println("Hello", keys)
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
	// Output:
	// Alice 2
	// Bob 3
	// Cecil 1
}
