package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubbleSort(values []int) []int {
	fmt.Printf("len to sort %d %v \n\r", len(values), values)
	for i := 0; i < len(values)-1; i++ {
		swap := false
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				tmp := values[j]
				values[j] = values[j+1]
				values[j+1] = tmp
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	return values
}

func main() {
	var value int
	var err error

	inStore := make([]int, 0, 3)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Print your value ")
		scanner.Scan()
		val := scanner.Text()
		val = strings.ToLower(val)
		if value, err = strconv.Atoi(val); err != nil {
			if val == "x" {
				break
			}
			fmt.Println("not an integer")
		} else {
			inStore = append(inStore, value)
			inStore = bubbleSort(inStore)
			fmt.Println("you put a new int value " + val)
			fmt.Printf("your slice values is %v len %d \n\r", inStore, len(inStore))
		}
	}
}
