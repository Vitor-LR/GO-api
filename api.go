package main

import (
	"fmt"
	"log"
	"net/http"

)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func main() {

	router := http.NewServeMux()
	router.HandleFunc("POST /users/{userID}", func(w http.ResponseWriter, r *http.Request){
		userID := r.PathValue("userID")
		w.Write([]byte("<h1>Hello "+userID+"</h1>"))
	})

	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	log.Printf("Server has Started %s", ":8080")

	err := server.ListenAndServe()
	if err != nil{
		fmt.Println(err.Error())
	}

}
