package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/routes"
	"github.com/kaasikodes/e-commerce-go/services"
)

type ApiServer struct {
	db *sql.DB
	addr string
}

func NewApiServer(db *sql.DB, addr string) *ApiServer {
	return &ApiServer{
		db: db,
		addr: addr,
	}
}

func (s *ApiServer) Start() error {
	router:= mux.NewRouter()
	// define subrouter to ensure api can easily be upgraded
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// define services/repos
	categoryRepo := services.NewCategoryRepository(s.db)
	tokenRepo := services.NewTokenRepository(s.db)
	userRepo := services.NewUserRepository(s.db)

	// define routes and map them to controllers
	routes.NewAuthRoutes(userRepo, tokenRepo).RegisterAuthRoutes(subrouter)
	routes.NewCategoryRoutes(categoryRepo).RegisterCategoryRoutes(subrouter)

	log.Println("Listening on ...", s.addr)
	return http.ListenAndServe(s.addr, router)
}