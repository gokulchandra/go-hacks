package main

import (
	"flag"
	"net/http"
)

func main() {
	port := flag.String("p", "7878", "Port to serve files" )
	dir := flag.String("d", "/", "Directory to serve")

	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))

	http.ListenAndServe(":"+*port, nil)


}