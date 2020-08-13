package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/msawatzky75/maintenence-log/server/endpoints"
	"github.com/msawatzky75/maintenence-log/server/graph/generated"
	"github.com/msawatzky75/maintenence-log/server/graph/model"
	graph "github.com/msawatzky75/maintenence-log/server/graph/resolvers"

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
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
			json.NewEncoder(w).Encode(map[string]string{
				"error": err,
			})
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
	})
	corsMiddleware := cors(os.Getenv("CORS_ORIGIN"))

	http.Handle("/graphiql", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", corsMiddleware.Handler(jwtMiddleware.Handler(srv)))
	http.Handle("/api/user", corsMiddleware.Handler(endpoints.User(db)))
	http.Handle("/api/auth/login", corsMiddleware.Handler(endpoints.Login(db)))
	http.Handle("/api/auth/refresh", corsMiddleware.Handler(endpoints.Refresh(db)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type CorsMiddleware struct{}

func cors(cors string) CorsMiddleware {
	return CorsMiddleware{}
}

func (*CorsMiddleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_ORIGIN"))

		switch r.Method {
		case "OPTIONS":
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.WriteHeader(http.StatusNoContent)
		}

		h.ServeHTTP(w, r)
	})
}
