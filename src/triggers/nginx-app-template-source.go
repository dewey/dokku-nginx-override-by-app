package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	fmt.Println("argsWithProg: ", argsWithProg)
	fmt.Println("triggered smoke-test-plugin from: nginx-app-template-source 2")
}
