package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello HttpRouter")
	})

	server := http.Server{
		Handler: router,
		Addr:    "Localhost:3000",
	}

	server.ListenAndServe()

}
