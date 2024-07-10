package routers

import (
	"log"
	"net/http"
	"time"

	"github.com/akposieyefa/open-ai/config"
	"github.com/akposieyefa/open-ai/core/handlers"
	"github.com/akposieyefa/open-ai/core/middlewares"
	response "github.com/akposieyefa/open-ai/pkg/http"
	"github.com/gorilla/mux"
)

func ApiRoutes() {

	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1").Subrouter()

	userRoute := api.PathPrefix("/users").Subrouter()
	settingsRoute := api.PathPrefix("/users").Subrouter()
	authRoute := api.PathPrefix("/auth").Subrouter()
	helpersRoute := api.PathPrefix("/helpers").Subrouter()

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response.WriteSuccessJson("Welcome to my simple ai api with chatgpt", map[string]string{}, true, w, http.StatusOK)
	}).Methods("GET")

	authRoute.HandleFunc("/login", handlers.Login).Methods("POST")
	authRoute.Use(middlewares.JwtVerify)
	authRoute.HandleFunc("/profile", handlers.GetProfile).Methods("GET")

	api.HandleFunc("/chat", handlers.ChatCompletionHandler).Methods("POST")

	userRoute.HandleFunc("/create", handlers.CreateUser).Methods("POST")

	helpersRoute.Use(middlewares.JwtVerify)
	helpersRoute.HandleFunc("/banks", handlers.GetAllBanks).Methods("GET")

	settingsRoute.HandleFunc("/change/password", handlers.CreateUser).Methods("POST")

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
