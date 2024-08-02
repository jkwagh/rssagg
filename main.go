package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	 portString := os.Getenv("PORT")
	 if portString == "" {
		log.Fatal("PORT is not found in the environment")
	 }

	 router := chi.NewRouter()

	 //Telling our server to send extra http headers in our responses to tell browsers we allow for the below
	 router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https//*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	 }))

	 //new router to mount to the /v1 path
	 v1Router := chi.NewRouter()
	 //handle only fires on GET requests
	 v1Router.Get("/healthz", handlerReadiness)
	 v1Router.Get("/err", handlerErr)

	 router.Mount("/v1", v1Router)

	 srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	 }

	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe()
	 if err != nil {
		log.Fatal(err)
	 }
}
