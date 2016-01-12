package main

import (
	"fmt"
)

var a, b int = 15, 6
var f float64

func main() {
	f = float64(a / b)
	fmt.Print(f)
}
