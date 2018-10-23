package main

import (
	"fmt"
	"net/http"
	"time"
	"api"
)

func main() {
	fmt.Print("Toolcat Start Running ... " + time.Now().String())
	mux := http.NewServeMux()

	//静态文件
	fileServer := http.FileServer(http.Dir("./public/files/"))
	http.StripPrefix("/static/", fileServer)

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/guid/new", api.Guid)

	server := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	server.ListenAndServe()
}
