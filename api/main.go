package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello Go!!</h1>")
}

func main() {
	str := "bye"
	num := 123
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8000", nil)
	fmt.Println(str)
	fmt.Println(num)
}
