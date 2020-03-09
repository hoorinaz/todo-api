package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", viewTodo)
	http.HandleFunc("/edit", editTodo)
	http.ListenAndServe(":8080", nil)

}

func viewTodo(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "View Hoorie Nazati")
}

func editTodo(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Edit Hoorie Nazati")

}
