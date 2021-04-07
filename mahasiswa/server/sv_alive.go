package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/satriaa14/test-api/mahasiswa/util/logging"
	"github.com/satriaa14/test-api/mahasiswa/util/setup"
)

func (rw *Service) Alive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setup.SetupCorsResponse(&w, r)
		log.Println(fmt.Sprintf("%+v", logging.LogRequestClient(w, r)))

		if r.URL.Path == root && r.Method == http.MethodGet {
			w.Write([]byte("Hi, Diah. I'm Alive :)"))
			return
		}
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
