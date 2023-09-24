package apiserver_internal

import (
	"net/http"

	index_function "github.com/I999m/ranking-api/inernal/api-server/func/index"
)

func AddRoutes(serverMux *http.ServeMux) {
	// ルーティングの設定

	serverMux.HandleFunc("/", index_function.Index())
}
