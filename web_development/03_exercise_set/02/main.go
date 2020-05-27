// Serve the files in the "starting-files" folder
// Use "http.FileServer"
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./images"))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "index.html")
}
