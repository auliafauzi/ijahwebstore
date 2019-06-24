package main

import (
	"flag"
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kataras/iris"
	"ijahwebstore/logger"
	// "ijahwebstore/middleware"
	"ijahwebstore/rest/webstore"
	// "os"
)

func main() {
	env := flag.String("env", "./Configuration/.env", "--env local|test|production")
	serverStart := flag.Bool("start", false, "-start true")
	flag.Parse()

	err := godotenv.Load(*env)
	if err != nil {
		logger.Warning.Println("Error loading .env file")
	}


	if *serverStart {
		fmt.Println("Web Store API Started..")
		app := iris.Default()

		crs := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
			AllowCredentials: true,
			AllowedHeaders:   []string{"x-access-token", "content-type"},
			AllowedMethods:   []string{"GET", "POST", "HEAD", "PATCH", "DELETE", "OPTIONS"},
		})

		// user.Register(app, "/user", crs, middleware.TokenValidation)
		webstore.Register(app,"/webstore", crs)

		app.Run(iris.Addr("0.0.0.0:8888"))
	}
}
