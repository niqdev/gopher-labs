package myhttp

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go

func StartServer() {
	log.Println("TODO not implemented")

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatalf("error %s", err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("new root request\n")
	io.WriteString(w, "ROOT\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("new hello request\n")
	io.WriteString(w, "HELLO\n")
}
