package main

import (
	"fmt"
	"log"
	"net/http"
)
func helloWorld(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method!="GET"{
		http.Error(w,"metod not supported",http.StatusNotFound)
		return
	}
	// fmt.Fprintf(w,"hello!")
	w.Write([]byte("hello world"))
}
func formHandler(w http.ResponseWriter,r *http.Request)  {
	if err := r.ParseForm();err!=nil{
		fmt.Fprintf(w,"PARSEFORM %V",err)
		return
	}
	fmt.Fprintf(w,"post request successful")
	fmt.Fprintf(w,"name : %s\n",r.FormValue("name"))
}
func main() {
	fileserver := http.FileServer((http.Dir("./static")))
	http.Handle("/",fileserver)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloWorld)
	fmt.Println("starting server at port 8080")
	if err := http.ListenAndServe(":8000",nil) ; err!=nil{
		log.Fatal(err)
	}
}