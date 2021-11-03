package main

import (
	"cgin/engine"
	"fmt"
	"net/http"
)

func main() {

	eng := engine.NewEngine()
	eng.GET("/index", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "WRL.PATH=%s\n", req.URL.Path)
	})

	http.ListenAndServe(":8080", eng)
}
