package main

import "fmt"

// Global variables
var (
	name             = "Jeth"
	firstName string = "James"
	lastName  string
)

// ========== COnstants ==============
const version = 2
func main() {
	fmt.Println(version)

	// ====== Local Variables and their declaration==============
	var version int
	version = 10
	fmt.Println(version)

	// version := 1 
	// otherVersion := "Buzz"
	// anotherVersion := 10.1


}
