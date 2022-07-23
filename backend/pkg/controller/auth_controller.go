package controller

import (
	"encoding/json"
	"net/http"
	"ruang-arah/backend/config"
	"ruang-arah/backend/helper"
	"ruang-arah/backend/model/web"
	"ruang-arah/backend/pkg/service"
	"time"
)

type AuthControllerImpl struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		authService: authService,
	}
}

func (c *AuthControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest web.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})
		return
	}

	err = loginRequest.ValidateLogin()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: http.StatusText(http.StatusUnprocessableEntity),
			Data:    err.Error(),
		})
		return
	}

	userLoginResponse, err := c.authService.GenerateToken(loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "SUCCESS",
		Data:    userLoginResponse,
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   userLoginResponse.Token,
		Expires: time.Now().Add(config.TOKEN_EXPIRES),
		Path:    "/",
	})

	helper.WriteToResponseBody(w, webResponse)
}

func (c *AuthControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	var registerRequest web.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})
		return
	}

	err = registerRequest.ValidateRegister()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: http.StatusText(http.StatusUnprocessableEntity),
			Data:    err.Error(),
		})
		return
	}

	registerResponse, err := c.authService.Register(
		registerRequest, registerRequest.Email,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(web.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    err.Error(),
		})
		return
	}

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "SUCCESS",
		Data:    registerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
