package engine

import (
	"fmt"
	"net/http"
)

type Handler func(w http.ResponseWriter, req *http.Request)

type Engine struct {
	getRoutes  map[string]Handler
	postRoutes map[string]Handler
}

func (engine *Engine) GET(path string, handler Handler) {
	engine.getRoutes[path] = handler
}

func (engine *Engine) POST(path string, handler Handler) {
	engine.postRoutes[path] = handler
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	var handler Handler
	var ok bool
	if req.Method == "GET" {
		handler, ok = engine.getRoutes[req.URL.Path]
		if !ok {
			return
		}
	} else if req.Method == "POST" {
		handler, ok = engine.getRoutes[req.URL.Path]
		if !ok {
			return
		}
	} else {
		fmt.Println("else method")
	}

	handler(w, req)
}

func NewEngine() *Engine {
	var engine Engine
	engine.getRoutes = make(map[string]Handler)
	engine.postRoutes = make(map[string]Handler)
	return &engine
}
