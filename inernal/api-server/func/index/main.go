package index_function

import "net/http"

func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Index"))
	}
}
