package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("mosquito bites"))
	fmt.Println("This is a string\nand there is a newline\nand another here!")
	fmt.Println("Tabs\tall\tover\tthe\tplace!!!")
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf("stringggg"))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
}
