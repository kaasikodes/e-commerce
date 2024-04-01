package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)


type Middleware func(http.HandlerFunc, ) http.HandlerFunc
func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.HandlerFunc,) http.HandlerFunc {
		for i := len(middlewares) -1 ;  i >= 0; i-- {

			next = middlewares[i](next)


		}
		return next
	}

}

func RequireAuthMiddleware(userRepo types.UserRepository) Middleware {
	return func(next http.HandlerFunc ) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request)  {
			tokenStr, err := utils.GetAccessTokenFromRequest(r)
			if err != nil {
				utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
				return
			}
			// validate jwt 
			token, err := utils.ValidateJWT(tokenStr)
			
			if err != nil {
				utils.WriteError(w, http.StatusUnauthorized, "Authorization Error", []error{err})
				return
			}
			if !token.Valid {
				err = fmt.Errorf("invalid token")
				utils.WriteError(w, http.StatusUnauthorized, "Authorization Error", []error{err})
				return
			}
	
			claims := token.Claims.(jwt.MapClaims)
			userIdMapKey := constants.JWTUserIdMapKey

			userId := claims[string(userIdMapKey)].(string)
	
			user, err := userRepo.RetrieveUserByID(userId)
			fmt.Println("user from auth middleware", user)
			if err != nil {
				utils.WriteError(w, http.StatusBadRequest, "User not found!", []error{err})
				return
			}
			ctx := context.WithValue(r.Context(), constants.JWTAuthUserContextKey, user)
			r = r.WithContext(ctx)
			next(w, r)
		}
	
	
	}

}

func RequirePermissionsMiddleware() {

}

func LoggerMiddleware() {

}
