package exception

import (
	"encoding/json"
	"net/http"
	"ruang-arah/backend/model/web"
)

func GET(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(web.WebResponse{
				Code:    http.StatusMethodNotAllowed,
				Message: "Need GET Method!",
				Data:    nil,
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func POST(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(web.WebResponse{
				Code:    http.StatusMethodNotAllowed,
				Message: "Need POST Method!",
				Data:    nil,
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func PUT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(web.WebResponse{
				Code:    http.StatusMethodNotAllowed,
				Message: "Need PUT Method!",
				Data:    nil,
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func DELETE(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(web.WebResponse{
				Code:    http.StatusMethodNotAllowed,
				Message: "Need DELETE Method!",
				Data:    nil,
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
