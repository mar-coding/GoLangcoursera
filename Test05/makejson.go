package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
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

func prettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func main() {
	m := make(map[string]string)
	name := prompt("Enter a name:")
	address := prompt("Enter a address:")
	m["name"] = name
	m["address"] = address
	barr, _ := json.Marshal(m)
	output, _ := prettyString(string(barr))
	print("-------------JSON-liked Output-------------\n")
	fmt.Println(output)
}
