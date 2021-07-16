package main

import "fmt"

func countSwaps(a []int32) {
	var count int
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			change := a[i]
			a[i] = a[i+1]
			a[i+1] = change
			count++
			i = -1
		}
	}

	fmt.Printf("Array is sorted in %v swaps.\n", count)
	fmt.Printf("First Element: %v\n", a[0])
	fmt.Printf("Last Element: %v\n", a[len(a)-1])
}

func main() {
	// countSwaps([]int32{1, 2, 3})    //Array is sorted in 0 swaps.\n First Element: 1\n Last Element: 3
	// countSwaps([]int32{3, 2, 1})    //Array is sorted in 3 swaps.\n First Element: 1\n Last Element: 3
	countSwaps([]int32{4, 2, 3, 1}) //Array is sorted in 5 swaps.\n First Element: 1\n Last Element: 4
}

/*
QUESTION ORIGINAL
package main

import (
    "fmt"
    "bufio"
    "io"
    "os"
    "strconv"
    "strings"
)


 //Complete the 'countSwaps' function below.
 //The function accepts INTEGER_ARRAY a as parameter.


 func countSwaps(a []int32) {

	var count int
	for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
					change := a[i]
					a[i] = a[i+1]
					a[i+1] = change
					count++
					i = -1
			}
	}

	fmt.Printf("Array is sorted in %v swaps.\n", count)
	fmt.Printf("First Element: %v\n", a[0])
	fmt.Printf("Last Element: %v\n", a[len(a)-1])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
			aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
			checkError(err)
			aItem := int32(aItemTemp)
			a = append(a, aItem)
	}

	countSwaps(a)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
			return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
			panic(err)
	}
}

*/
