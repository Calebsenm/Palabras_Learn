package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
		
)

func getRoot(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got / request\n")
	io.WriteString(w,"this is my website! \n")
}

// rutas del servidor
func getHello(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/index"{
		http.Error(w, "404 not founde", http.StatusNotFound)
		return
	}

}


func main(){

	fileServer := http.FileServer(http.Dir("../src"))
	http.Handle("/",fileServer)

	http.HandleFunc("/index",getHello)



	err := http.ListenAndServe(":3333", nil)
	fmt.Printf("Starting at the port :3333")

	if errors.Is(err,http.ErrServerClosed){
		fmt.Printf("Server closed\n")

	}else if err != nil{
		fmt.Printf("Error Starting server: %s\n",err)
		os.Exit(1)
		
	}
}
