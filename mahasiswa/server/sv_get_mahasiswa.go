package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/satriaa14/test-api/mahasiswa/middleware"
	"github.com/satriaa14/test-api/mahasiswa/util/logging"
	"github.com/satriaa14/test-api/mahasiswa/util/pattern"
	"github.com/satriaa14/test-api/mahasiswa/util/setup"
)

func (rw *Service) GetMahasiswa() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println(fmt.Sprintf("%+v", logging.LogRequestClient(w, r)))

		setup.SetupCorsResponse(&w, r)

		if pattern.DataURL(r.URL.Path)[4] != getall {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		// Middle ware
		err := middleware.MiddlewareJwtAuth(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		if r.Method != http.MethodGet {
			http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
			return
		}

		// TODO
		resp, err := rw.repo.SQLiteReadWriter.GetMahasiswa()
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		data, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		w.Write(data)
	}
}
