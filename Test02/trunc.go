package main

import "fmt"

func main() {
	var floatNum float32
	fmt.Print("Enter a floating point number:")
	fmt.Scan(&floatNum)
	fmt.Printf("%.0f", float32(int(floatNum)))
}
