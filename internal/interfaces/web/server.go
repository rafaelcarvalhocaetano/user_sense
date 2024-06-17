package web

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(handlers chi.Router) {

	handlers.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("<h1>welcome to the deepbot</h1>"))
		if err != nil {
			return
		}
	})
	handlers.Mount("/api", handlers)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%v", os.Getenv("WEB_PORT")),
		Handler:        handlers,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	url := os.Getenv("ENDPOINT")
	log.Println(fmt.Sprintf("xodo run in: %v" + url))
	log.Fatal(server.ListenAndServe())
}
