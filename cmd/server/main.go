package main

import (
	"fmt"
	"net/http"

	"github.com/GiovaniGitHub/multithreading/configs"
	_ "github.com/GiovaniGitHub/multithreading/docs"
	"github.com/GiovaniGitHub/multithreading/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Desafio 2.0 - Multithreading
// @version         1.0
// @description     Fullcycle Pós Go Expert Go Expert

// @termsOfService  http://swagger.io/terms/

// @contact.name   Giovani Angelo
// @contact.email  giovani.angelo@gmail.com

// @host      localhost:8000
// @BasePath  /
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/cep", func(r chi.Router) {
		r.Get("/{cep}", handlers.GetAddress)
	})
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:"+configs.WebServerPort+"/docs/doc.json")))
	http.ListenAndServe(fmt.Sprintf(":%s", configs.WebServerPort), r)
}
