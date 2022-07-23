package service

import (
	"ruang-arah/backend/helper"
	"ruang-arah/backend/model/domain"
	"ruang-arah/backend/model/web"
	"ruang-arah/backend/pkg/repository"
)

type FrontendServiceImpl struct {
	authRepository     repository.AuthRepository
	frontendRepository repository.FrontendRepository
	backendRepository  repository.BackendRepository
}

func NewFrontendService(
	authRepo repository.AuthRepository,
	frontendRepo repository.FrontendRepository,
	backendRepo repository.BackendRepository,
) *FrontendServiceImpl {
	return &FrontendServiceImpl{
		authRepository:     authRepo,
		frontendRepository: frontendRepo,
		backendRepository:  backendRepo,
	}
}

func (service *FrontendServiceImpl) GetProgrammingLanguages() ([]web.ProgrammingLanguageResponse, error) {
	programmingLanguages, err := service.frontendRepository.FindAllProgrammingLanguages()
	helper.PanicIfError(err)

	return helper.ToProgrammingLanguageResponses(programmingLanguages), nil
}

func (service *FrontendServiceImpl) GetQuestionByProgrammingLanguageIdWithPagination(
	programmingLanguageId, page, limit int32) ([]web.QuestionResponse, error) {
	questions, err := service.frontendRepository.FindAllQuestionByProgrammingLanguageIdWithPagination(
		programmingLanguageId, page, limit,
	)
	helper.PanicIfError(err)

	programmingLanguageDomain, err := service.backendRepository.FindProgrammingLanguageById(
		programmingLanguageId,
	)
	helper.PanicIfError(err)

	var incorrectAnswer domain.IncorrectAnswerDomain
	for _, question := range questions {
		answerDomain, err := service.backendRepository.FindIncorrectAnswersByQuestionId(
			question.Id,
		)
		helper.PanicIfError(err)

		incorrectAnswer = answerDomain
	}

	return helper.ToQuestionResponses(
		questions, incorrectAnswer, programmingLanguageDomain,
	), nil
}

func (service *FrontendServiceImpl) SubmitAnswersAttempts(userId int32, request web.AnswerAttemptRequest) (web.AnswersAttemptsResponse, error) {
	var (
		answersAttempt []domain.AnswerAttemptDomain
		level          domain.LevelDomain
		levelName      string
	)

	for _, answer := range request.Answers {
		answerAttemptDomain := []domain.AnswerAttemptDomain{
			{UserId: userId, QuestionId: answer.QuestionId, Answer: answer.Answer},
		}

		answersAttempt = append(answersAttempt, answerAttemptDomain...)

	}

	_, err := service.frontendRepository.SaveAnswersAttempts(userId, answersAttempt)
	helper.PanicIfError(err)

	total, err := service.frontendRepository.FindTotalAnswersAttemptsByUserId(userId)
	helper.PanicIfError(err)

	if total >= 75 {
		levelName = "Intermediate"
	} else {
		levelName = "Beginner"
	}

	if levelName != "" {
		levelDomain, err := service.frontendRepository.FindLevelByName(levelName)
		helper.PanicIfError(err)

		level = levelDomain
	}

	user, err := service.authRepository.FindUserById(userId)
	helper.PanicIfError(err)

	return helper.ToAnswersAttemptsResponse(user, level, total), nil
}

func (service *FrontendServiceImpl) GetRecommendationByLevelId(levelId int32) (web.RecommendationResponse, error) {
	recommendation, err := service.frontendRepository.FindRecommendationByLevelId(levelId)
	helper.PanicIfError(err)

	return helper.ToRecommendationResponse(recommendation), nil
}
