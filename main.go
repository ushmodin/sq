package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type rq struct {
	Data   []byte
	Header http.Header
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use: " + os.Args[0] + " [port]")
		os.Exit(1)
	}

	queue := make(chan rq, 1)
	lock := make(chan bool, 1)
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			<-lock
			queue <- rq{Data: data, Header: r.Header}
		} else if r.Method == "GET" {
			lock <- true
			rq := <-queue

			for hn, hv := range rq.Header {
				for _, vv := range hv {
					w.Header().Add(hn, vv)
				}
			}

			_, err := w.Write(rq.Data)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		} else {
			http.Error(w, "Method not implemented", 500)
		}
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+os.Args[1], nil))
}
