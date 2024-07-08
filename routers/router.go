package routers

import (
	"log"
	"net/http"
	"time"

	"github.com/akposieyefa/open-ai/config"
	"github.com/akposieyefa/open-ai/core/handlers"
	response "github.com/akposieyefa/open-ai/pkg/http"
	"github.com/gorilla/mux"
)

func ApiRoutes() {

	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response.WriteSuccessJson("Welcome to my simple ai api with chatgpt", map[string]string{}, true, w, http.StatusOK)
	}).Methods("GET")

	api.HandleFunc("/chat", handlers.ChatCompletionHandler).Methods("POST")

	srv := &http.Server{
		Handler:        router,
		Addr:           config.Envs.APP_PORT,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		IdleTimeout:    120 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
