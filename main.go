package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type MyJson struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: ./program num1 num2\n")
		return
	}
	code, _ := strconv.Atoi(os.Args[1])
	val1 := 1337
	val2, _ := strconv.ParseUint(os.Args[2], 10, 64)
	val3, _ := strconv.Atoi(os.Args[3])
	var myObj MyJson
	jsonStr := `{"a":1337, "b":0}`
	json.Unmarshal([]byte(jsonStr), &myObj)
	switch code {
	case 1:
		result := val1 / val3
		fmt.Println(result)
	case 2:
		tmp := int(val2)
		result := val1 / tmp
		fmt.Println(result)
	case 3:
		result := val1 / int(val2)
		fmt.Println(result)
	case 4:
		tmp := int(val2)
		if tmp == 0 {
			fmt.Println(tmp)
		}
	case 5:
		tmp := int(val2)
		if tmp != 0 {
			boom := val1 / tmp
			fmt.Println(boom)
		} else {
			fmt.Println("Ok")
		}
	case 6:
		a, b := myObj.A, myObj.B
		if b == 0 {
			boom := a / b
			fmt.Println(boom)
		}
	case 7:
		a, b := myObj.A, myObj.B
		fmt.Printf("%d\n", a / b)
	}
	return
}
