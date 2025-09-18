package main

import "fmt"

func main() {
	
	fmt.Println("Starting")
	//panic("An unrecoverable error happened")
	//fmt.Println("Ending") // this line will never be executed
	result := safeDivide(10,0)
	fmt.Println("Result: ", result)
	fmt.Println("Ending")

}

func safeDivide(a,b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverred from panic: ", r)
		}
	}()
	return a/b
}