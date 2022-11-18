package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func prompt(toprint string) string {
	if toprint == "" {
		toprint = "Enter Integer Number or \"X\": "
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(toprint)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func isNum(intNum string) bool {
	if _, err := strconv.Atoi(intNum); err != nil {
		return false
	}
	return true
}

func main() {
	var sli = make([]int, 0, 3)
	var tempSli = make([]string, 0, 3)
	var multiNum = false
	for {
		input := prompt("")
		if strings.Compare(strings.ToLower(input), "x") == 0 {
			break
		}
		if strings.Contains(input, " ") {
			tempSli = strings.Split(input, " ")
			multiNum = true
		}
		if multiNum {
			for i := 0; i < len(tempSli); i++ {
				if !isNum(tempSli[i]) {
					continue
				}
				num, _ := strconv.Atoi(tempSli[i])
				sli = append(sli, num)
				sort.Sort(sort.IntSlice(sli))
			}
			break
		}
		if isNum(input) {
			num, _ := strconv.Atoi(input)
			sli = append(sli, num)
			sort.Sort(sort.IntSlice(sli))
			fmt.Println(sli)
		} else {
			continue
		}
	}
	fmt.Printf("final list is: %v", sli)
}
