package main

import "fmt"

func main() {
	// fmt.Print("Its variables topics")
	// var userName string = "ahmad"
	// fmt.Println(userName)
	// fmt.Printf("userName is of type : %T \n ", userName)
	// var isLoggenIn bool = true
	// fmt.Println(isLoggenIn)
	// fmt.Printf("userName is of type : %T \n ", isLoggenIn)

	// default values and some  alias as well

	var anotherVariable int // Here is this fun fact that it does not put any garbage value but put 0 according to here
	fmt.Println(anotherVariable)
	fmt.Printf("The Type of this Language is %T \n", anotherVariable)

	// Implicit Type
	var website = "learncodeonline.in" // here the lexer play the role for type identification
	fmt.Println(website)
	// but you cant assign any other value to this

	// No var style code
	website1 := "www.golang.com"
	fmt.Println(website1)

}
