package apiserver_internal

import (
	"log"
	"net/http"
)

type Options struct {
	AccessControlAllowOrigin      string
	AccessControlAllowHeaders     string
	AccessControlAllowCredentials string
	AccessControlAllowMethods     string
}

func builder(serverMux *http.ServeMux, options *Options) http.Handler {

	return http.HandlerFunc(
		// リクエスト毎に実行される処理を定義
		func(w http.ResponseWriter, r *http.Request) {
			// ログ出力
			log.Printf("Request: %s %s", r.Method, r.URL.Path)

			// CORS対応
			w.Header().Set("Access-Control-Allow-Origin", options.AccessControlAllowOrigin)
			w.Header().Set("Access-Control-Allow-Headers", options.AccessControlAllowHeaders)
			w.Header().Set("Access-Control-Allow-Credentials", options.AccessControlAllowCredentials)
			w.Header().Set("Access-Control-Allow-Methods", options.AccessControlAllowMethods)

			serverMux.ServeHTTP(w, r)
		})
}
