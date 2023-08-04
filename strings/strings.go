package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var a_string = "some fucking string"

	fmt.Println("has some prefix: " + strconv.FormatBool(strings.HasPrefix(a_string, "some")))
	fmt.Println("to upper case:" + strings.ToUpper(a_string))
	fmt.Println("it contains fucking: " + strconv.FormatBool(strings.Contains(a_string, "fucking")))

	s := []string{"this is", a_string}

	fmt.Println(strings.Join(s, " "))
	fmt.Printf("%q\n", strings.Split("this is a goddamn long string", " "))
	fmt.Println(strings.ReplaceAll(a_string, " ", "_"))
}
