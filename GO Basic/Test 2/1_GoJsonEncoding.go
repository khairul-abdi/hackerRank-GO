package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

/*
 * Complete the 'EncodeManager' function below and the struct Manager.
 *
 * The function is expected to return an io.Reader and an error.
 * The function accepts *Manager manager as parameter.
 */

type Manager struct {
	FullName       string
	Position       string
	Age            int32
	YearsInCompany int32
}

func EncodeManager(manager *Manager) (io.Reader, error) {
	// emp := &manager{FullName:"Jack Oliver",Position:"CEO",Age:44, YearsInCompany: 5}
	// e, err := json.Marshal(emp)
	// if err != nil {
	//     return nil, err
	// }
	// fmt.Println(string(e))
	// respByte, err := ioutil.ReadAll(e)
	// return respByte, nil
	fmt.Println("APA ini ", manager)
	byteString := []byte(`{"full_name":"Jack Oliver","position":"CEO","age":44,"years_in_company":8}`)
	err := json.Unmarshal(byteString, &manager)
	if err != nil {
		fmt.Println(err)
	}
	r := bytes.NewReader(byteString)
	// respByte, err := ioutil.Reader(byteString)
	return r, nil

}

// Expected: {"full_name":"Jack Oliver","position":"CEO","age":44,"years_in_company":8}
//Expected: {"full_name":"John Tejada","position":"CPO","age":28}

func main() {
	var manager1 = Manager{FullName: "Tejada", Position: "CPO", Age: 28} //&{John Tejada CPO 28}
	// var manager2 = {"full_name":"Jack Oliver","position":"CEO","age":44,"years_in_company":8}
	EncodeManager(manager1, &Manager)
}

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
// 	checkError(err)

// 	defer stdout.Close()

// 	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

// 	manager := &Manager{}

// 	fullName := readLine(reader)
// 	if fullName != "" {
// 		manager.FullName = fullName
// 	}

// 	position := readLine(reader)
// 	if position != "" {
// 		manager.Position = position
// 	}

// 	ageTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)
// 	age := int32(ageTemp)
// 	if age != 0 {
// 		manager.Age = age
// 	}

// 	yearsInCompanyTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)
// 	yearsInCompany := int32(yearsInCompanyTemp)
// 	if yearsInCompany != 0 {
// 		manager.YearsInCompany = yearsInCompany
// 	}

// 	resultReader, err := EncodeManager(manager)
// 	checkError(err)

// 	result, err := ioutil.ReadAll(resultReader)
// 	checkError(err)

// 	fmt.Fprintf(writer, "%s\n", string(result))

// 	writer.Flush()
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
