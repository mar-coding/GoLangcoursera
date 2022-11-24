
package main

import (
		"fmt"
		"io/ioutil"
		"strings"
		"os"
	)
type Name struct {
	fname string
	lname string
}
func main() {
	var sli []Name
	var fpath string
	// Prompt file path
	fmt.Printf("Please the file path: ")
	fmt.Scanln(&fpath)
	
	// Read the file
	content, e := ioutil.ReadFile(fpath)
	if e != nil {os.Exit(1)}
	lines := strings.Split(string(content), "\n")

	// Process each line
	for _, v := range lines {
		flnames := strings.Split(v, " ")
		n := Name{fname: flnames[0],  lname: flnames[1]}
		sli = append(sli, n)
	}

	// Print every line (fname lname)
	for _, v := range sli {
		fmt.Println(v.fname, v.lname)
	}
	
		
}
