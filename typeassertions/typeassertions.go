package main

import "fmt"

func main() {
	var value interface {} = 42
	num := value.(int)
	fmt.Println(num)

	var value2 interface {} = 33.42
	num2, ok := value2.(float64)
	if ok {
		fmt.Println(num2)
	} else {
		fmt.Println("value2 is not float64")
	}

	describe(num2)
}

func describe(value interface {}) {
	switch v := value.(type) {
	case int:
		fmt.Println("int", v)
	case float64:
		fmt.Println("float64", v)
	default:
		fmt.Println("unkown", v)		
	}
	
}