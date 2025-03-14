package main

import (
	"fmt"
	router "wv-api-publicaciones-go/router"
)

func main() {
	name := "Idealista Microservices with Go"
	fmt.Println("Whitevega", name)

	runServer()
}

func runServer() {
	router.Init()
}
