package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	var half, marathon float64
	half, marathon = 13.1, 26.2
	hero := "Deena Kastor"
	hero2 := "paula radcliffe"

	fmt.Println(math.Floor(half))
	fmt.Println(strings.Title(hero2))
	fmt.Println("This is a string\nand there is a newline\nand another here!")
	fmt.Println("Tabs\tall\tover\tthe\tplace!!!\t", hero)
	fmt.Println(reflect.TypeOf(marathon))
	fmt.Println(reflect.TypeOf("stringggg"))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
}
