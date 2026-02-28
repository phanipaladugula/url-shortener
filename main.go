package main

import(
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Pong !! The server is ready to kick")
}

func main(){
	http.HandleFunc("/ping",pingHandler)

	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080",nil)
}