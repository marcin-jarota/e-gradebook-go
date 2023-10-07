package main

import (
	"e-student/internal/adapters/repository"
	"e-student/internal/adapters/service"
	"e-student/internal/adapters/storage"
	"e-student/internal/app/domain"
	"e-student/internal/common"
	transport "e-student/internal/http"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := common.NewConfig()

	conn, err := gorm.Open(postgres.Open(cfg.Dsn), nil)

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	userRepo := repository.NewGormUserRepository(conn)
	markRepo := repository.NewGormMarkRepository(conn)
	studentRepo := repository.NewGormStudentRepository(conn)
	subjectRepo := repository.NewGormSubjectRepository(conn)

	tpls := template.Must(template.ParseFiles("./web/templates/pages/login.html"))

	engine := html.New("./web/templates", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	storage := storage.NewMemoryStorage()
	authService := service.NewAuthService(userRepo, storage, cfg)
	studentSrvc := service.NewStudentService(studentRepo)
	studentHandler := transport.NewStudentHandler(studentSrvc)
	userHandler := transport.NewUserHandler(authService)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3001,http://localhost:3002",
	}))

	app.Get("/", userHandler.GetLogin)

	app.Post("/", userHandler.PostLogin)

	app.Get("/delete", func(c *fiber.Ctx) error {
		storage.Delete(c.Query("id"))
		return c.JSON(fiber.Map{"message": fmt.Sprintf("Deleted session for user %s", c.Query("id"))})
	})

	// app.Get("/students2", userHandler.AuthMiddleware(), studentHandler.GetAllStudents)

	appGroup := app.Group("/app", userHandler.AuthMiddleware())

	appGroup.Get("/home", func(c *fiber.Ctx) error {
		return c.Render("pages/home", nil, "layouts/main")
	})

	appGroup.Get("/students", studentHandler.GetAllStudents)

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	type response struct {
	// 		Data any `json:"data"`
	// 	}
	//
	// 	user, err := userRepo.GetOne(12)
	//
	// 	if err != nil {
	// 		w.WriteHeader(500)
	// 		parsed, _ := json.Marshal(response{
	// 			Data: err.Error(),
	// 		})
	//
	// 		w.Header().Set("Content-Type", "application/json")
	// 		w.Write(parsed)
	// 		return
	// 	}
	//
	// 	parsed, _ := json.Marshal(response{
	// 		Data: user,
	// 	})
	//
	// 	fmt.Println(parsed)
	// 	w.WriteHeader(200)
	//
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(parsed)
	// })

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		// tpl, err := template.ParseFiles("./cmd/web/index.html")

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// err = tpl.Execute(w, nil)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		err := tpls.ExecuteTemplate(w, "login.html", nil)

		if err != nil {
			log.Fatal(err)
		}
	})

	r.Get("/test1", func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.ParseFiles("./cmd/web/tst.html")

		user, err := userRepo.GetAll()
		if err != nil {
			log.Fatal(err)
		}

		students, err := studentRepo.GetAll()
		if err != nil {
			log.Fatal(err)
		}

		err = tpl.Execute(w, map[string]any{
			"User":     user,
			"Students": students,
		})
		if err != nil {
			log.Fatal(err)
		}
	})

	r.Get("/students12", func(w http.ResponseWriter, r *http.Request) {
		type response struct {
			Error bool              `json:"error"`
			Data  []*domain.Student `json:"data"`
		}
		responseError := false

		students, err := studentRepo.GetAll()
		// w.Header().Add("Content-Type", "application/json")

		if err != nil {
			responseError = true
		}

		jsonResponse, _ := json.Marshal(response{
			Data:  students,
			Error: responseError,
		})

		w.Write(jsonResponse)
	})

	r.Get("/subjects", func(w http.ResponseWriter, r *http.Request) {
		type response struct {
			Error bool              `json:"error"`
			Data  []*domain.Subject `json:"data"`
		}
		responseError := false

		subjects, err := subjectRepo.GetAll()
		w.Header().Add("Content-Type", "application/json")

		if err != nil {
			responseError = true
		}

		jsonResponse, _ := json.Marshal(response{
			Data:  subjects,
			Error: responseError,
		})

		w.Write(jsonResponse)
	})

	r.Post("/subject", func(w http.ResponseWriter, r *http.Request) {
		var subject domain.Subject
		err := json.NewDecoder(r.Body).Decode(&subject)
		w.Header().Add("Content-Type", "application/json")

		if err != nil {
			log.Fatal(err)
			return
		}

		err = subjectRepo.AddSubject(&subject)

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

	r.Post("/mark", func(w http.ResponseWriter, r *http.Request) {
		var mark domain.Mark

		err := json.NewDecoder(r.Body).Decode(&mark)

		w.Header().Add("Content-Type", "application/json")

		if err != nil {
			handleErr(err, w)
			return
		}

		students, err := studentRepo.GetAll()
		if err != nil {
			handleErr(err, w)
			return
		}

		subjects, err := subjectRepo.GetAll()
		if err != nil {
			handleErr(err, w)
			return
		}

		mark.StudentID = students[0].ID
		mark.SubjectID = subjects[0].ID

		err = markRepo.AddMark(&mark)

		if err != nil {
			handleErr(err, w)
			return
		}

		w.WriteHeader(http.StatusOK)

		response, _ := json.Marshal(struct {
			Ok bool `json:"ok"`
		}{Ok: true})

		w.Write(response)

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

	r.Post("/student", func(w http.ResponseWriter, r *http.Request) {
		var user domain.User
		var student domain.Student

		err := json.NewDecoder(r.Body).Decode(&user)
		w.Header().Add("Content-Type", "application/json")

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response, _ := json.Marshal(struct {
				Error string `json:"error"`
			}{Error: err.Error()})

			w.Write(response)
			return
		}

		fmt.Println("USER PASWOD ", user.Password, user.Name)
		student.User = user

		err = studentRepo.AddStudent(&student)

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

	log.Printf("Listening on port http://localhost:%s", cfg.Port)

	fs := http.FileServer(http.Dir("./web/dist"))

	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	app.Static("/static", "./web/dist")
	app.Listen(fmt.Sprintf(":%s", cfg.Port))
	// http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r)

}

func handleErr(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	log.Panic(err)
	response, _ := json.Marshal(struct {
		Error string `json:"error"`
	}{Error: err.Error()})

	w.Write(response)
}
