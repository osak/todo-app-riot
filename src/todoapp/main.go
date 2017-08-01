package main

import (
	"net/http"
	"io"
	"os"
	"encoding/json"
	"fmt"
	"strings"
)

type Todo struct {
	Title     string `json:"title"`
	Completed bool `json:"completed"`
}

func serveDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s\n", r.URL.Path)
	f, err := os.Open("static/" + r.URL.Path)
	defer f.Close()

	if err != nil {
		w.WriteHeader(404)
		return
	}
	if strings.HasSuffix(r.URL.Path, ".html") {
		w.Header().Set("Content-Type", "text/html")
	} else if strings.HasSuffix(r.URL.Path, ".js") {
		w.Header().Set("Content-Type", "text/javascript")
	}
	io.Copy(w, f)
}

func listTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode([]Todo{
		Todo{Title: "a", Completed: false, },
		Todo{Title: "b", Completed: true, },
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", serveDefault)
	mux.HandleFunc("/todos", listTodo)
	http.ListenAndServe("localhost:8080", mux)
}
