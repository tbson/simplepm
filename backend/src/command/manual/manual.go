package main

import (
	"fmt"
	"os"
)

func main() {
	arg0 := os.Args[1:][0]
	fmt.Println(arg0)
}
