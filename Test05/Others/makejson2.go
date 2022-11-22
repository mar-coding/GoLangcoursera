package main

import (
    "fmt"
    "encoding/json"
    "os"
    "bufio"
)

func main() {
	m := make( map[string]string )
	//
	fmt.Println( "Please enter your name:" )
	
	scan_name := bufio.NewScanner( os.Stdin )
	scan_name.Scan()
	m["name"] = scan_name.Text()
	//
	fmt.Println( "Then input your adress:" )
	
	scan_addr := bufio.NewScanner( os.Stdin )
	scan_addr.Scan()
	m["adress"] = scan_addr.Text()
	//
    b, err := json.Marshal(m)
    if err != nil {
		fmt.Println( "Encoding faild" )
		
    } else {
        fmt.Println( "Encoded data : ", b )
		fmt.Println( "Decoded data : ", string(b) )
		
	}

}