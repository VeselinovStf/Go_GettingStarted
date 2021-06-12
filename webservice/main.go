package main

import (
	"Go_GettingStarted/webservice/controllers"
	"net/http"
)


func main(){
	controllers.RegisterControllers()
	http.ListenAndServe(":3000",nil)
}