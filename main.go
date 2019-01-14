package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
)

type rq struct {
	Data   []byte
	Header http.Header
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("FCGI-Server. Implement blocking producer-consumer pattern.\nUse: " + os.Args[0] + " [port]")
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
			ioutil.ReadAll(r.Body)
			http.Error(w, "Method not implemented", 500)
		}
	}

	l, err := net.Listen("tcp", ":"+os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if err := fcgi.Serve(l, http.HandlerFunc(handler)); err != nil {
		log.Fatal(err)
	}
}
