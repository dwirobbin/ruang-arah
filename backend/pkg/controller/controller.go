package controller

import (
	"net/http"
	"ruang-arah/backend/pkg/service"
)

type AuthController interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type BackendController interface {
	CreateQuestion(w http.ResponseWriter, r *http.Request)
	GetQuestions(w http.ResponseWriter, r *http.Request)
	GetQuestionById(w http.ResponseWriter, r *http.Request)
	UpdateQuestion(w http.ResponseWriter, r *http.Request)
	DeleteQuestion(w http.ResponseWriter, r *http.Request)
}

type FrontendController interface {
	GetProgrammingLanguages(w http.ResponseWriter, r *http.Request)
	GetQuestionByProgrammingLanguageIdWithPagination(w http.ResponseWriter, r *http.Request)
	SubmitAnswersAttempts(w http.ResponseWriter, r *http.Request)
	GetRecommendationByLevelId(w http.ResponseWriter, r *http.Request)
}

type Controller struct {
	AuthController
	BackendController
	FrontendController
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		AuthController:     NewAuthController(service.AuthService),
		BackendController:  NewBackendController(service.BackendService),
		FrontendController: NewFrontendController(service.FrontendService),
	}
}
