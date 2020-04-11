package main

import (
	"fmt"
	"strconv"
)

func main() {

	var decimal int64

	fmt.Print(10)
	fmt.Scanln(&decimal)

	output := strconv.FormatInt(decimal, 2)

	fmt.Print("Output ", output)

}
