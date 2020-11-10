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

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const defaultPort = "4000"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"))

	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		log.Println(connectionString)
		log.Fatal(err)
	}

	db.AutoMigrate(
		&model.User{},
		&model.UserPreference{},
		&model.Vehicle{},
		&model.FuelLog{},
		&model.MaintenanceLog{},
		&model.OilChangeLog{})
	defer db.Close()

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

	http.Handle("/graphiql", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", corsMiddleware.Handler(jwtMiddleware.Handler(srv), "POST"))
	http.Handle("/api/user", corsMiddleware.Handler(loginEndpoint.LoginHandler(), "GET, POST"))
	http.Handle("/api/signup", corsMiddleware.Handler(signupEndpoint.SignupHandler(), "PUT"))
	http.Handle("/api/auth/login", corsMiddleware.Handler(loginEndpoint.LoginHandler(), "POST"))
	http.Handle("/api/auth/logout", corsMiddleware.Handler(loginEndpoint.LogoutHandler(), "POST"))
	http.Handle("/api/auth/refresh", corsMiddleware.Handler(loginEndpoint.RefreshHandler(), "POST"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
