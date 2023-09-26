package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("hello")
	fmt.Println("world")
}


func processFile(fileName string) error {
	file, err := os.Open(fileName)

	if err != nil {
		return err
	}
	
	defer file.Close()


	//process your file
	return nil
}