package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	consoleReader := bufio.NewReader(os.Stdin)
	print("Please enter a string:")
	input, _ := consoleReader.ReadString('\n')
	str = strings.ToLower(input)
	str = strings.Join(strings.Fields(str), "")

	aChar := strings.Count(str, "a")
	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && aChar > 0 {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}
