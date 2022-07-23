package service

import (
	"ruang-arah/backend/model/web"
	"ruang-arah/backend/pkg/repository"
)

type AuthService interface {
	GenerateToken(request web.LoginRequest) (web.LoginResponse, error)
	Register(request web.RegisterRequest, email string) (web.RegisterResponse, error)
}

type BackendService interface {
	CreateQuestion(request web.QuestionRequest) (web.QuestionResponse, error)
	GetQuestions() ([]web.QuestionListResponse, error)
	GetQuestionById(questionId int32) (web.QuestionResponse, error)
	UpdateQuestion(questionId int32, request web.QuestionRequest) (web.QuestionResponse, error)
	DeleteQuestion(questionId int32) (bool, error)
}

type FrontendService interface {
	GetProgrammingLanguages() ([]web.ProgrammingLanguageResponse, error)
	GetQuestionByProgrammingLanguageIdWithPagination(programmingLanguageId, page, limit int32) ([]web.QuestionResponse, error)
	SubmitAnswersAttempts(userId int32, request web.AnswerAttemptRequest) (web.AnswersAttemptsResponse, error)
	GetRecommendationByLevelId(levelId int32) (web.RecommendationResponse, error)
}

type Service struct {
	AuthService
	BackendService
	FrontendService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthService:    NewAuthService(repo.AuthRepository),
		BackendService: NewBackendService(repo.BackendRepository),
		FrontendService: NewFrontendService(
			repo.AuthRepository, repo.FrontendRepository, repo.BackendRepository,
		),
	}
}
