package main

import (
	"mwitter-backend/src/rest"
)

func main() {

	r := rest.RunAPI()

	r.Run(":8080")
}
