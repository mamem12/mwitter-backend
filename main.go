package main

import (
	"mwitter-backend/src/rest"
)

func main() {

	r := rest.RunAPI()

	// r.HandleContext(&s)

	r.Run(":8080")
}
