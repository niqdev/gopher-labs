package myhttp

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go

func StartServer() {
	log.Println("listening on port 3333")

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/hello", getHello)

	server := &http.Server{
		Addr:              ":3333",
		ReadHeaderTimeout: 3 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("error %s", err)
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("new hello request\n")
	io.WriteString(w, "HELLO\n")
}
