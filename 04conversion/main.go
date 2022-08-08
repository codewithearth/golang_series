package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our coding world")
	fmt.Println("Please rate our work between 1 and 5 ")

	// read input value
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating : ", input)

	//convert string into float
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	//handle error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add 1 in your rating :  ", numRating+1)
	}
}
