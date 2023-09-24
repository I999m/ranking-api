package apiserver_internal

import (
	"errors"
	"fmt"
	"net/http"
)

func Run(port int) {
	fmt.Println(port)

	serverMux := http.NewServeMux()

	AddRoutes(serverMux)

	handler := builder(serverMux, &Options{
		AccessControlAllowOrigin:      "*",
		AccessControlAllowHeaders:     "Content-Type",
		AccessControlAllowCredentials: "true",
		AccessControlAllowMethods:     "GET,PUT,POST,DELETE,UPDATE,OPTIONS",
	})

	addr := fmt.Sprintf(":%d", port)

	fmt.Println("Server started on" + addr)

	server := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("error running http server: %s\n", err)
		}
	}

}
