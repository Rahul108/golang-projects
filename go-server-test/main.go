package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err:= r.ParseForm(); err!=nil {
		fmt.Fprintf(w, "ParseFrom() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request success")
	name:= r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
	}

	fmt.Fprint(w, "Hello!")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting Server at port 8000")
	if err:= http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}


}