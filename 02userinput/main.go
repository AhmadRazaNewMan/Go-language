package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input file"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your understanding in go langaugae:")

	//Comma Ok or Comma Err syntacs in go language there is no try and catch the go treats the error also as an error

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating", input)  // The Type of the input is String So now Will now handle it 

}
