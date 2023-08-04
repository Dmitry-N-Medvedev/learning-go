package main

import (
	"fmt"
	"strings"
)

func main() {

	var a_string = "some fucking string"

	fmt.Println(strings.HasPrefix(a_string, "some"))
}
