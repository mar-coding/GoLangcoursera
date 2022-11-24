package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func read_lines(address string, line_size int) []string {
	// if line_size equals to 0 it use default value
	// and the default value is 64k
	var output []string
	file, err := os.Open(address)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if line_size != -1 {
		maxCapacity := line_size
		buf := make([]byte, maxCapacity)
		scanner.Buffer(buf, maxCapacity)
	}
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
	return output
}

func prompt(toprint string) string {
	if toprint == "" {
		toprint = "Enter input: "
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(toprint)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	var names []Name
	file_name := prompt("Enter input name:")
	lines := read_lines(file_name, -1)
	for _, line := range lines {
		temp := strings.Split(line, " ")
		var name_temp Name
		name_temp.fname = temp[0]
		name_temp.lname = temp[1]
		names = append(names, name_temp)
	}
	for _, item := range names {
		fmt.Printf("%s , %s\n", item.fname, item.lname)
	}

}
