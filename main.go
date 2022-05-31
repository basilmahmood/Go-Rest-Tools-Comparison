package main

import (
	"fmt"
	"os"
)

var data = make(map[string]interface{})
var id = 1

func main() {
	tool := os.Args[1]

	if tool == "net" {
		netAPI()
	} else if tool == "gin" {
		ginAPI()
	} else if tool == "echo" {
		echoAPI()
	} else {
		fmt.Errorf("tool must be one of 'net', 'gin', or 'echo")
	}

}
