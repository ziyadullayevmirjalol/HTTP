package main
import (
	"fmt"
	"net/http"
)
func main() {
	port := ":8080"
	fmt.Println("Server working on port :8080")

	http.HandleFunc("/", HandleRoot)
	
	err := http.ListenAndServe(port, nil) 
	if err != nil {
		fmt.Println("Something Wrong...")
	}
}
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			fmt.Println("GET method")
			fmt.Fprint(w, "Sent")
		case "POST":
			fmt.Println("POST method")
			fmt.Fprint(w, "Sent")
		case "PUT":
			fmt.Println("PUT method")
			fmt.Fprint(w, "Sent")
		case "DELETE":
			fmt.Println("Delete method")
			fmt.Fprint(w, "Sent")
	}
}
