package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func Swap(num *[]int, idx int) {
	temp := (*num)[idx]
	(*num)[idx] = (*num)[idx+1]
	(*num)[idx+1] = temp
}

func BubbleSort(num *[]int) {
	length := len(*num)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if (*num)[j] > (*num)[j+1] {
				Swap(num, j)
			}
		}
	}
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
	var nums []int
	input := prompt("Enter 10 number:")
	temp := strings.Split(input, " ")
	for _, n := range temp {
		t, _ := strconv.Atoi(n)
		nums = append(nums, t)
	}
	BubbleSort(&nums)
	fmt.Printf("%v", nums)
}
