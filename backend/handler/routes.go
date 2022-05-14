package handler

import (
	"net/http"

	"github.com/tris-tux/go-lis/backend/db"
)

func InitRoutes(postgres *db.Postgres) *http.ServeMux {
	taskHandler := &taskHandler{
		postgres: postgres,
		static:   &db.Static{},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/static", taskHandler.GetStatic)
	mux.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, cache-control")

		switch r.Method {
		case http.MethodOptions:
			w.Write([]byte("allowed"))
		case http.MethodGet:
			taskHandler.getAllTask(w, r)
		case http.MethodPost:
			taskHandler.insertTask(w, r)
		case http.MethodPut:
			taskHandler.updateTask(w, r)
		case http.MethodDelete:
			taskHandler.deleteTask(w, r)
		default:
			responseError(w, http.StatusNotFound, "")
		}
	})

	return mux
}
