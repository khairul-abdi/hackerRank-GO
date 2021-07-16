package main

import (
	"fmt"
	"strings"
)

func ModifyString(str string) string {
	var res []string
	// var res string

	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == "0" || string(str[i]) == "1" || string(str[i]) == "2" || string(str[i]) == "3" || string(str[i]) == "4" || string(str[i]) == "5" || string(str[i]) == "6" || string(str[i]) == "7" || string(str[i]) == "8" || string(str[i]) == "9" {
			continue
		} else if string(str[i]) != " " {
			res = append(res, string(str[i]))
			// res += string(str[i])
		}
	}
	// fmt.Println(res)
	return strings.Join(res[:], "")
	// fmt.Println(str)
	// return str
}

func main() {
	fmt.Println(ModifyString("6d6WKTByBJKi27wqE5rJm2AECodXVBZvDdOaXa6LkKkohk0eUBUTu4pd31eak6WUaEhPXtzYs9UJB4PpewBYpGRiezH0BOxZ5iO8xuzzIuLcf7n102tpb04RanIfk"))
	fmt.Println(ModifyString("another")) //rehtona
}
