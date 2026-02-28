package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ShortenRequest struct{
	URL string `json:"url"`
}

type ShortenResponse struct{
	ShortURL string `json:"short_url"`
}

func pingHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println(w,"Pong!! The server is ready to kick")
}
func shortenHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w,"Use Post", http.StatusMethodNotAllowed)
		return
	}
	var req ShortenRequest
	json.NewDecoder(r.Body).Decode(&req)

	var id uint64
	err := dbPool.QueryRow(r.Context(),
	"INSERT INTO urls (original_url,short_code) VALUES ($1,$2) RETURNING id",
	req.URL,"").Scan(&id)
	
	if err != nil{
		http.Error(w,"Database error",http.StatusInternalServerError)
	}

	shortCode := Encode(id)
	_,err = dbPool.Exec(r.Context(),"UPDATE urls SET short_code = $1 WHERE id = $2", shortCode,id)

	if err != nil{
		http.Error(w,"Failed to save short code", http.StatusInternalServerError)
		return
	}

	resp := ShortenResponse{ShortURL: "http://localhost:8080/"+shortCode}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(resp)

}
func main(){
	initDB()

	defer dbPool.Close()

	http.HandleFunc("/ping",pingHandler)
	http.HandleFunc("/shorten",shortenHandler)

	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080",nil)
}