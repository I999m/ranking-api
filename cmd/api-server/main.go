package main

import (
	"flag"

	apiserver_internal "github.com/I999m/ranking-api/inernal/api-server"
)

func main() {
	// ポート番号を格納する変数を定義
	var port int

	// コマンドライン引数 port 定義
	flag.IntVar(&port, "port", 9000, "listen port for the api server")

	// コマンドライン引数のパース
	flag.Parse()

	// サーバーの起動
	apiserver_internal.Run(port)
}
