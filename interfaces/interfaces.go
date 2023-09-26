package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() string {
	return "Woof!"
}

func main() {
	var mySpeaker Speaker = &Dog{
		Name : "Buddy",
	}
	
	fmt.Println(mySpeaker.Speak())
	introduce(mySpeaker)

	printAnything("Hello World.")
	printAnything(42)
}

func introduce(s Speaker) {
	fmt.Println(s.Speak())
}

//The empty interface can be used when the actual type is unknown or when a function should accept any type.
func printAnything(a interface{}) {
	fmt.Println(a)
}