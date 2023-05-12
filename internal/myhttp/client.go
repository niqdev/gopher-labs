package myhttp

import (
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

// TODO http tests

func SimpleHttpRequest() {
	resp, err := http.Get("http://localhost:8080/ip")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func RetryHttpRequest() {
	resp, err := retryablehttp.Get("http://localhost:8080/status/504")
	if err != nil {
		panic(err)
	}
	fmt.Println("status", resp.Status)
}
