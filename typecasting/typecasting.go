package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	var f float64 = float64(i)
	fmt.Println("Integer value:", i)
	fmt.Println("Float value:", f)

	var j int = 10
	var k float64 = 3.5
	sum := float64(j) + k
	fmt.Println("Sum:", sum)

	str := strconv.Itoa(i)
	fmt.Println("String: ", str)

	num,err := strconv.Atoi(str)
	if err == nil {
		fmt.Println("int: ", num)
	} else {
		fmt.Println("error: ", err)
	}
}