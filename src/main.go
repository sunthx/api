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
	//fileServer := http.FileServer(http.Dir("./public/files/"))
	//mux.Handle("/files/", http.StripPrefix("/files/", fileServer))

	mux.HandleFunc("/guid/new", api.Guid)
	mux.HandleFunc("/file/upload",api.UploadFile)
	mux.HandleFunc("/file/download/",api.DownloadFile)

	server := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	server.ListenAndServe()
}
