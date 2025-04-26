package main

import (
	"fmt"
	"log"
	"net/http"

)
func hellohandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found" , http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error( w , "methood not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"Hello go learners")
	
}
func formhandler(w http.ResponseWriter, r *http.Request){
	if err:= r.ParseForm(); err != nil{
		fmt.Fprintf(w,"Parseform() err: %v",err)
		return
	}
	fmt.Fprintf(w, "post request successful\n");
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name =%s ",name)
	fmt.Fprintf(w,"Address =%s ",address)
}
func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/form", formhandler)
	fmt.Print("started the port 8080")
	if err:= http.ListenAndServe(":8080",nil);err != nil{
		log.Fatal(err)
	}
}