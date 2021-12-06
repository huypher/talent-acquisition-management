package infra

import "net/http"

func NewRestService(httpHandler http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: httpHandler,
	}
}
