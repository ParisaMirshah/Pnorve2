// main.go

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func sumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	num1 := r.FormValue("num1")
	num2 := r.FormValue("num2")

	sum, err := calculateSum(num1, num2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "The sum is: %s", sum)
}

func calculateSum(num1, num2 string) (string, error) {
	n1, err := parseFloat(num1)
	if err != nil {
		return "", err
	}

	n2, err := parseFloat(num2)
	if err != nil {
		return "", err
	}

	sum := n1 + n2
	return fmt.Sprintf("%.2f", sum), nil
}

func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func main() {
	http.HandleFunc("/sum", sumHandler)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
