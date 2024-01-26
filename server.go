package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, I'm %s. I'm %s years old.", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	fmt.Fprintf(w, "My family: %s", string(data))
}
