package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/up", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
