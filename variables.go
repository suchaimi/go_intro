package main

import "fmt"

func main() {
	var x string = "Hello, World"
	fmt.Println(x)

	var s string
	s = "First"
	fmt.Println(s)
	s = "Second"
	fmt.Println(s)

	holderName := "Rocky Gerung"
	fmt.Println("The holder name is", holderName)
	dogName := "Pitbull"
	fmt.Println("The dog name is", dogName)
	dogName = "Siberian"
	fmt.Println("The dog name is", dogName)
}
