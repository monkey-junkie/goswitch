package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: ./program code val1 val2 val3\n")
		return
	}
	code, _ := strconv.Atoi(os.Args[1])
	val1 := 1337
	val2, _ := strconv.ParseUint(os.Args[2], 10, 64)
	val3, _ := strconv.Atoi(os.Args[3])

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
	}
	return
}
