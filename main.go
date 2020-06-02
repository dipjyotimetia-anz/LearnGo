package main

import (
	"learnGo/controllers"
	"net/http"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000",nil)
}
