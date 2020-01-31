package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		res := x
		x, y = y, x+y // swap
		return res
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin) // reader
	fmt.Print("input: ")
	input, _ := reader.ReadString('\n') // get input
	input = strings.TrimSpace(input)    // remove \n from string
	int, _ := strconv.Atoi(input)       // str to int

	f := fibonacci()
	for i := 0; i < int; i++ {
		fmt.Print(f(), ",")
	}
}
