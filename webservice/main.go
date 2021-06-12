package main

import (
	"Go_GettingStarted/webservice/models"
	"fmt"
)


func main(){
	u := models.User{
		Id: 2,
		FirstName: "John",
		LastName: "Dow",
	}

	fmt.Println(u)
}