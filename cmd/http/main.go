package main

import (
	"echo-crate/internal/server"
	"fmt"
)

func main() {

	server := server.New()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
