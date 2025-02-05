package main

import (
	"fmt"
	"modules/types"
	"modules/utils"
) 


func main()  {
	// === getNumber =====
	// number := getNumber()

	//  ==== User struct type =====
	user := types.User {
		Username: utils.GetUsename(),
		Age: utils.GetAge(),
	}

	fmt.Printf("The user is: %+v ", user)

}