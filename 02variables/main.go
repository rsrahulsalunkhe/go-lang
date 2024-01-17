package main

import "fmt"

const LoginToken string = "ghabhgjd" //public

func main() {
	var username string = "Rahul"
	fmt.Println(username)
	fmt.Printf("Variabe is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variabe is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variabe is of type: %T \n", smallVal)

	var samllfloat float64 = 255.66985472
	fmt.Println(samllfloat)
	fmt.Printf("Variabe is of type: %T \n", samllfloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variabe is of type: %T \n", anotherVariable)

	// implicit type
	var website = "rahul salunkhe"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}
