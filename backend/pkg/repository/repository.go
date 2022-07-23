package repository

import (
	"database/sql"
	"ruang-arah/backend/model/domain"
)

type AuthRepository interface {
	FindUser(email, password string) (domain.UserDomain, error)
	FindUserById(id int32) (domain.UserDomain, error)
	Save(user domain.UserDomain, email string) (domain.UserDomain, error)
}

type BackendRepository interface {
	FindProgrammingLanguageByName(name string) (domain.ProgrammingLanguageDomain, error)
	FindProgrammingLanguageById(id int32) (domain.ProgrammingLanguageDomain, error)
	FindProgrammingLanguageByQuestions(questions []domain.QuestionDomain) ([]domain.ProgrammingLanguageDomain, error)
	FindQuestionById(id int32) (domain.QuestionDomain, error)
	FindIncorrectAnswersByQuestionId(questionId int32) (domain.IncorrectAnswerDomain, error)
	SaveQuestion(question domain.QuestionDomain, incorrectaAnswer domain.IncorrectAnswerDomain) (domain.QuestionDomain, error)
	FindAllQuestions() ([]domain.QuestionDomain, error)
	UpdateQuestion(question domain.QuestionDomain, incorrectaAnswer domain.IncorrectAnswerDomain) (domain.QuestionDomain, error)
	DeleteQuestion(questionId int32) (bool, error)
}

type FrontendRepository interface {
	FindAllProgrammingLanguages() ([]domain.ProgrammingLanguageDomain, error)
	FindAllQuestionByProgrammingLanguageIdWithPagination(programmingLanguageId, page, limit int32) ([]domain.QuestionDomain, error)
	SaveAnswersAttempts(userId int32, request []domain.AnswerAttemptDomain) (bool, error)
	FindTotalAnswersAttemptsByUserId(userId int32) (int32, error)
	FindLevelByName(levelName string) (domain.LevelDomain, error)
	FindRecommendationByLevelId(levelId int32) (domain.RecommendationDomain, error)
}

type Repository struct {
	AuthRepository
	BackendRepository
	FrontendRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		AuthRepository:     NewAuthRepository(db),
		BackendRepository:  NewBackendRepository(db),
		FrontendRepository: NewFrontendRepository(db),
	}
}
