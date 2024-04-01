package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/models"
	"github.com/kaasikodes/e-commerce-go/types"
	"github.com/kaasikodes/e-commerce-go/utils"
)

type AuthController struct {
	userRepo types.UserRepository
	tokenRepo types.TokenRepository
}

func NewAuthController(userRepo types.UserRepository, tokenRepo types.TokenRepository) *AuthController{
	return &AuthController{
		userRepo: userRepo,
		tokenRepo: tokenRepo,
	}
}


// Forgot Pwd
func (h *AuthController) ForgotPwdHandler(w http.ResponseWriter, r *http.Request)  {
	repo := h.userRepo
	var payload types.ForgotPwdInput
	
	if err:= utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
		
	}
	errParsed := utils.ValidatePayload(payload)
	if len(errParsed) > 0{
	
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, errParsed)
		return
	}
	user, err := repo.RetrieveUserByEmail(payload.Email)
	if err != nil {
		err  = fmt.Errorf("user with email %s not found", payload.Email)
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}

	// create pwd reset token
	token, err := h.tokenRepo.CreatePasswordResetToken(types.CreateTokenInput{
		Email: user.Email,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	// send pwd reset email
	err = sendPasswordResetEmail(user.Email, token.Token)
	
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Hello, a password reset link has been sent to your email",  user)
		
}
// Register User
func (h *AuthController) RegisterUserHandler(w http.ResponseWriter, r *http.Request)  {
	repo := h.userRepo
	var payload types.AddUserInput
	
	if err:= utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
		
	}
	errParsed := utils.ValidatePayload(payload)
	if len(errParsed) > 0{
	
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, errParsed)
		return
	}
	// encrypt password
	encryptedPassword, err := utils.EncryptPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	payload.Password = encryptedPassword
	// add user
	user, err := repo.AddUser(payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
	}
	// create verification token
	token, err := h.tokenRepo.CreateVerificationToken(types.CreateTokenInput{
		Email: user.Email,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	// send verification email
	err = sendVerificationEmail(user.Email, token.Token)
	
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	utils.WriteJson(w, http.StatusOK, "Congratulations, your account has been created, please verify your email!",  user)
		
}
func sendPasswordResetEmail(userEmail string, token string)error {
	err := utils.SendMail([]string{userEmail}, "Reset Password", fmt.Sprintf("Please reset your password by clicking on the link: %s/auth/reset-password?token=%s&email=%s", constants.FrontendUrl, token, userEmail))
	return err
}
func sendVerificationEmail(userEmail string, token string)error {
	err := utils.SendMail([]string{userEmail}, "Verify your account", fmt.Sprintf("Please verify your email by clicking on the link: %s/auth/verify?token=%s&email=%s", constants.FrontendUrl, token, userEmail))
	return err
}
// Reset Pwd
func (h *AuthController) ResetPwdrHandler(w http.ResponseWriter, r *http.Request)  {
	t := h.tokenRepo
	u := h.userRepo
	var payload types.ResetPwdInput
	
	if err:= utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
		
	}
	errParsed := utils.ValidatePayload(payload)
	if len(errParsed) > 0{
	
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, errParsed)
		return
	}
	
	token, err := t.RetrievePasswordResetToken(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
	}
	// verify token
	if payload.Token != token.Token {
		err = fmt.Errorf("invalid token")
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
	}
	// check token expiration
	if time.Now().After(token.ExpiresAt) {
		err = fmt.Errorf("token expired")
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
	}
	// verify user
	user, err := u.VerifyUser(payload.Email)


	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	// delete token
	err = t.DeletePasswordResetToken(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	
	
	utils.WriteJson(w, http.StatusOK, "Password reset successful!",  user)
		
}
// Verify User
func (h *AuthController) VerifyUserHandler(w http.ResponseWriter, r *http.Request)  {
	t := h.tokenRepo
	u := h.userRepo
	var payload types.VerifyTokenInput
	
	if err:= utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
		
	}
	errParsed := utils.ValidatePayload(payload)
	if len(errParsed) > 0{
	
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, errParsed)
		return
	}
	
	token, err := t.RetrieveVerificationToken(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
	}
	// verify token
	if payload.Token != token.Token {
		err = fmt.Errorf("invalid token")
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
	}
	// check token expiration
	if time.Now().After(token.ExpiresAt) {
		err = fmt.Errorf("token expired")
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
	}
	// verify user
	user, err := u.VerifyUser(payload.Email)


	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	// delete token
	err = t.DeleteVerificationToken(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	
	
	utils.WriteJson(w, http.StatusOK, "User verified!",  user)
		
}

func  (h *AuthController) ensureUserEmailIsVerified(user models.User, payloadEmail string) (bool, error) {
	emailVerified := true
	if !user.EmailVerified {
		
		token, err := h.tokenRepo.RetrieveVerificationToken(payloadEmail)
		if err != nil {
			emailVerified = false
			return emailVerified, err
		}
		if  time.Now().After(token.ExpiresAt) {
			// delete token
			err = h.tokenRepo.DeleteVerificationToken(payloadEmail)
			if err != nil {
				emailVerified = false
				return emailVerified, err
			}
			// create new token
			token, err = h.tokenRepo.CreateVerificationToken(types.CreateTokenInput{
				Email: user.Email,
			})
			if err != nil {
				emailVerified = false
				return emailVerified, err
			}
			
		}
		
		err = sendVerificationEmail(user.Email, token.Token)
	
		if err != nil {
			emailVerified = false
			return emailVerified, err
		}
		err = fmt.Errorf("user not verified, please check your email, a verification link has been sent to your email")
		return emailVerified, err
	}
	return emailVerified, nil
}
// login user

func  (h *AuthController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserInput
	err:= utils.ParseJSON(r, &payload);
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
	}
	errParsed := utils.ValidatePayload(payload)
	if len(errParsed) > 0{
	
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, errParsed)
		return
	}
	// check if user exists
	user, err := h.userRepo.RetrieveUserByEmail(payload.Email)
	if err != nil {
		err = fmt.Errorf("user not found")
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
	}
	// check if password is correct
	isValidPassword := utils.CheckPasswordHash(payload.Password,user.Password)
	
	if !isValidPassword {
		err = fmt.Errorf("invalid password")
		utils.WriteError(w, http.StatusBadRequest, constants.MsgValidationError, []error{err})
		return
	}
	// check if the user has been 
	isEmailVerified, err := h.ensureUserEmailIsVerified(user, payload.Email)
	if err != nil || !isEmailVerified {
		utils.WriteError(w, http.StatusBadRequest, "User not verified!", []error{err})
		return
	}

	
	token, err := utils.CreateJWT(user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, constants.MsgInternalServerError, []error{err})
		return
	}
	authData := createAuthResponseData(user, token)
	utils.WriteJson(w, http.StatusOK, "User logged in successfully!", authData)

}


func createAuthResponseData(user models.User, token string) map[string]interface{} {
	return map[string]interface{}{
		"accessToken": token,
		"user": user,
	}
}