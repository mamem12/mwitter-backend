package main

import "mwitter-backend/src/rest"

func main() {
	// fmt.Println("Hello World")
	ginEngine := rest.RunAPI()
	ginEngine.Run(":8000")
}
