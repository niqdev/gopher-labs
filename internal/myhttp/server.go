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

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/status", getStatus)

	fs := http.FileServer(http.Dir("./internal/myhttp/public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/home/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/public/", http.StatusSeeOther)
	})

	server := &http.Server{
		Addr:              ":3333",
		ReadHeaderTimeout: 3 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("error %s", err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("[%s] new root request", r.Host))
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "ROOT")
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("new status request")
	io.WriteString(w, "OK")
}
