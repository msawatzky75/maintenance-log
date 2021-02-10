package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/msawatzky75/maintenance-log/server/endpoints"
	"github.com/msawatzky75/maintenance-log/server/graph/generated"
	"github.com/msawatzky75/maintenance-log/server/graph/model"
	graph "github.com/msawatzky75/maintenance-log/server/graph/resolvers"
	"github.com/msawatzky75/maintenance-log/server/middleware"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "4000"

func main() {
	err := godotenv.Load()
	if err != nil {
		// .env is supported but not required, you can pass in the environment via the system variables instead.
		log.Print("Error loading .env file")
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"))

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		i := 0
		for i < 10 {
			i++
			sqlDB, _ := db.DB()
			err := sqlDB.Ping()
			dbg(err)
			time.Sleep(time.Second * 2)
		}
	}

	if err != nil {
		dbgf("Error opening connection to database at %s:%s",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"))
		log.Fatal(err)
	}

	// Will not delete old fields, this will only create columns and tables that don't exist
	db.AutoMigrate(
		&model.User{},
		&model.UserPreference{},
		&model.Vehicle{},
		&model.FuelLog{},
		&model.MaintenanceLog{},
		&model.OilChangeLog{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))
	jwtMiddleware := middleware.Jwt{
		AccessTokenCookie:  "access_token",
		RefreshTokenCookie: "refresh_token",
		Secret:             []byte(os.Getenv("JWT_SECRET")),
	}
	corsMiddleware := middleware.Cors{Cors: os.Getenv("CORS_ORIGIN")}
	dbgf("CORS: %s", os.Getenv("CORS_ORIGIN"))
	loginEndpoint := endpoints.Login{
		DB:                 db,
		JWTSecret:          []byte(os.Getenv("JWT_SECRET")),
		AccessTokenCookie:  "access_token",
		AccessTokenLife:    time.Minute * 15,
		RefreshTokenCookie: "refresh_token",
		RefreshTokenLife:   time.Hour * 24 * 7,
		CookieDomain:       os.Getenv("HOST_DOMAIN"),
	}
	signupEndpoint := endpoints.Signup{
		DB: db,
	}

	// Don't want the playground in prod
	if os.Getenv("APP_ENV") == "DEVELOPMENT" {
		http.Handle("/graphiql", playground.Handler("GraphQL playground", "/graphql"))
		log.Printf("connect to http://localhost:%s/graphiql for GraphQL playground", port)
	}

	http.Handle("/graphql", corsMiddleware.Handler(jwtMiddleware.Handler(srv), "POST"))
	http.Handle("/api/user", corsMiddleware.Handler(loginEndpoint.LoginHandler(), "GET, POST"))
	http.Handle("/api/signup", corsMiddleware.Handler(signupEndpoint.SignupHandler(), "PUT"))
	http.Handle("/api/auth/login", corsMiddleware.Handler(loginEndpoint.LoginHandler(), "POST"))
	http.Handle("/api/auth/logout", corsMiddleware.Handler(loginEndpoint.LogoutHandler(), "POST"))
	http.Handle("/api/auth/refresh", corsMiddleware.Handler(loginEndpoint.RefreshHandler(), "POST"))

	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func dbg(items ...interface{}) {
	if os.Getenv("LOG_LEVEL") == "DEBUG" {
		log.Print(items...)
	}
}
func dbgf(template string, items ...interface{}) {
	if os.Getenv("LOG_LEVEL") == "DEBUG" {
		log.Printf(template, items...)
	}
}
