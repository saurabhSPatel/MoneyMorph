package main

import (
	"fmt"
	"net/http"

	"github.com/MoneyMorph/server"
)

func main() {
	http.HandleFunc("/convert", server.ConvertCurrency)
	port := ":8080"
	fmt.Println("Server is running on http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
