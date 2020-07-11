package main

import (
	"fmt"
)
import "github.com/xueyubingsen/idcardconv/vars"

func main() {
	var s1 string = vars.Mmm(100)
	fmt.Printf("I got %v", s1)
}
