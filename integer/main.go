package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i1 int = 38
	fmt.Printf("i1: %s\n", strconv.Itoa(i1))

	fmt.Printf("i1: %s\n", fmt.Sprintf("%d", i1))

}
