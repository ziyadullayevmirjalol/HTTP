
package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":8080"
	fmt.Println("Server working on port :8080")

	err := http.ListenAndServe(port, nil) 
	if err != nil {
		fmt.Println("Something Wrong...")
	}
}
