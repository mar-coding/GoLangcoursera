package main

import (
	"fmt"
	"strconv"
	"sort"
	"strings"
)

func main()  {
	s := make([]int, 0, 3)
	var userInput string
	fmt.Println("Provide:\n* an integer\n* 'X' to exit")
	for i := 0; i < 1;  {
 	 	fmt.Scan(&userInput)
		userInput = strings.ToLower(userInput)
		if userInput == "x" {
			i++;
		} else {
			el,err := strconv.Atoi(userInput)
			if err == nil {
				s = appendAndSort(s, el)
				fmt.Println(s)
			}
		}
	}
}

func appendAndSort(s []int, elment int) []int {
	s = append(s, elment)
	sort.Ints(s)
	return s
}