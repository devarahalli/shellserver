package shellserver

import (
	"log"
	"net/http"
)

func StartHttpServer(addr string) error {
	log.Fatal(http.ListenAndServe(addr, nil))
	return nil
}
