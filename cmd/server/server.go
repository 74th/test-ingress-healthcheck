package main

import (
	"fmt"
	"net/http"
	"os"
)

type sv struct {
}

func (s *sv) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	w.WriteHeader(200)
	w.Write([]byte(url.Path))
}

func main() {
	host := "0.0.0.0:8080"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}
	fmt.Print("start", host)
	err := http.ListenAndServe(host, &sv{})
	if err != nil {
		fmt.Print(err.Error())
	}
}
