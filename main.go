package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":80", "set http listen addr")
}

func main() {
	flag.Parse()

	if strings.HasPrefix(addr, ":") {
		fmt.Printf("================ START httpDump http://127.0.0.1%s ================\n", addr)
	} else {
		fmt.Printf("================ START httpDump http://%s ================\n", addr)
	}

	err := http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Printf("-----------[%d] [%s] recv from %s-----------\n", now.Unix(), now.Format("2006-01-02 15:04:05"), r.RemoteAddr)
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))

		w.Write(requestDump)
		r.Body.Close()
	}))
	if err != nil {
		fmt.Println(err)
	}
}
