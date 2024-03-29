package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaasikodes/e-commerce-go/controllers"
	"github.com/kaasikodes/e-commerce-go/types"
)

type AuthRoutes struct {
	userRepo types.UserRepository
	tokenRepo types.TokenRepository
}

func NewAuthRoutes(userRepo types.UserRepository, tokenRepo types.TokenRepository) *AuthRoutes {
	return &AuthRoutes{
		userRepo: userRepo,
		tokenRepo: tokenRepo,
	}
}

func (c *AuthRoutes) RegisterAuthRoutes (router *mux.Router){
	controller := controllers.NewAuthController(c.userRepo, c.tokenRepo)
	router.HandleFunc("/register", controller.RegisterUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/login", controller.LoginUser).Methods(http.MethodPost)
	router.HandleFunc("/forgot-password", controller.ForgotPwdHandler).Methods(http.MethodPost)
	router.HandleFunc("/reset-password", controller.ResetPwdrHandler).Methods(http.MethodPatch)
	router.HandleFunc("/verify-user", controller.VerifyUserHandler).Methods(http.MethodPost)
}