package main

import (
    "fmt"
 	"os"
 	"strings"
	 "bufio"
	 "encoding/json"
)
func main() {
	
	type Person struct {
		Fname string
		Lname string
	}
	var listPeople []Person
	var fileName string

	fmt.Println("Write the file name: ")
	fmt.Scan(&fileName)
	readFile, err := os.Open(fileName)
  
    if err != nil {
        fmt.Println("No such file")
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
    for fileScanner.Scan() {
		names := strings.Split(fileScanner.Text()," ")
		person := Person{Fname: names[0], Lname: names[1]}
		listPeople = append(listPeople, person)
    }
	fmt.Println("List people: ", listPeople)

	fmt.Println("List people JSON: ")
	for i, _ := range listPeople {
		u, _ := json.Marshal(listPeople[i])
		fmt.Println(string(u))
	}
    readFile.Close()
	
}