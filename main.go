package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		res := x
		x, y = y, x+y
		return res
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin) // reader
	fmt.Print("input: ")
	input, _ := reader.ReadString('\n') // get input
	input = strings.TrimSpace(input)    // remove \n from string
	int, _ := strconv.Atoi(input)

	f := fibonacci()
	for i := 0; i < int; i++ {
		fmt.Print(f(), ",")
	}
}
