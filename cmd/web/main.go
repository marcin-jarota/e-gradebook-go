package main

import (
	"e-student/internal/adapters/repository"
	"e-student/internal/domain"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	dsn := flag.String("dsn", os.Getenv("DSN"), "Connection DSN")
	port := flag.String("port", os.Getenv("PORT"), "Port to listen on")
	flag.Parse()

	conn, err := gorm.Open(postgres.Open(*dsn), nil)

	conn.AutoMigrate(&domain.User{})

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	userRepo := repository.NewUserRepository(conn)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		type response struct {
			Data any `json:"data"`
		}

		user, err := userRepo.GetOne(1)

		if err != nil {
			w.WriteHeader(500)
			parsed, _ := json.Marshal(response{
				Data: err.Error(),
			})

			w.Header().Set("Content-Type", "application/json")
			w.Write(parsed)
			return
		}

		parsed, _ := json.Marshal(response{
			Data: user,
		})

		fmt.Println(parsed)
		w.WriteHeader(200)

		w.Header().Set("Content-Type", "application/json")
		w.Write(parsed)
	})

	r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		var p domain.User

		err := json.NewDecoder(r.Body).Decode(&p)
		w.Header().Add("Content-Type", "application/json")

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response, _ := json.Marshal(struct {
				Error string `json:"error"`
			}{Error: err.Error()})

			w.Write(response)
			return
		}

		err = userRepo.AddUser(&p)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response, _ := json.Marshal(struct {
				Error string `json:"error"`
			}{Error: err.Error()})

			w.Write(response)
			return
		}

		w.WriteHeader(http.StatusOK)

		response, _ := json.Marshal(struct {
			Ok bool `json:"ok"`
		}{Ok: true})

		w.Write(response)

	})

	log.Printf("Listening on port http://localhost:%s", *port)

	http.ListenAndServe(fmt.Sprintf(":%s", *port), r)

}
