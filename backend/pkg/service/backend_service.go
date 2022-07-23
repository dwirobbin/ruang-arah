package service

import (
	"ruang-arah/backend/helper"
	"ruang-arah/backend/model/domain"
	"ruang-arah/backend/model/web"
	"ruang-arah/backend/pkg/repository"
	"time"
)

type BackendServiceImpl struct {
	backendRepository repository.BackendRepository
}

func NewBackendService(backendRepo repository.BackendRepository) *BackendServiceImpl {
	return &BackendServiceImpl{
		backendRepository: backendRepo,
	}
}

func (service *BackendServiceImpl) CreateQuestion(questionReq web.QuestionRequest) (web.QuestionResponse, error) {
	programmingLanguage, err := service.backendRepository.FindProgrammingLanguageByName(
		questionReq.ProgrammingLanguage,
	)
	helper.PanicIfError(err)

	questionDomain := domain.QuestionDomain{
		Question:              questionReq.Question,
		ProgrammingLanguageId: programmingLanguage.Id,
		CorrectAnswer:         questionReq.CorrectAnswer,
	}

	incorrectAnswerDomain := domain.IncorrectAnswerDomain{
		QuestionId: questionDomain.Id,
		OptionA:    questionReq.IncorrectOne,
		OptionB:    questionReq.IncorrectTwo,
	}

	questionDomainResponse, err := service.backendRepository.SaveQuestion(
		questionDomain, incorrectAnswerDomain,
	)
	helper.PanicIfError(err)

	return helper.ToQuestionResponse(
		programmingLanguage, questionDomainResponse, incorrectAnswerDomain,
	), nil
}

func (service *BackendServiceImpl) GetQuestions() ([]web.QuestionListResponse, error) {
	questions, err := service.backendRepository.FindAllQuestions()
	helper.PanicIfError(err)

	programmingLanguages, err := service.backendRepository.FindProgrammingLanguageByQuestions(questions)
	helper.PanicIfError(err)

	return helper.ToQuestionListResponses(questions, programmingLanguages), nil
}

func (service *BackendServiceImpl) GetQuestionById(questionId int32) (web.QuestionResponse, error) {
	question, err := service.backendRepository.FindQuestionById(questionId)
	helper.PanicIfError(err)

	programmingLanguage, err := service.backendRepository.FindProgrammingLanguageById(
		question.ProgrammingLanguageId,
	)
	helper.PanicIfError(err)

	incorrectAnswer, err := service.backendRepository.FindIncorrectAnswersByQuestionId(
		question.Id,
	)
	helper.PanicIfError(err)

	return helper.ToQuestionResponse(programmingLanguage, question, incorrectAnswer), nil
}

func (service *BackendServiceImpl) UpdateQuestion(questionId int32, questionReq web.QuestionRequest) (web.QuestionResponse, error) {
	questionResponse, err := service.backendRepository.FindQuestionById(questionId)
	helper.PanicIfError(err)

	var programmingLanguage domain.ProgrammingLanguageDomain
	var question domain.QuestionDomain
	if questionReq.ProgrammingLanguage != "" {
		progLangDomainRepo, err := service.backendRepository.FindProgrammingLanguageByName(
			questionReq.ProgrammingLanguage,
		)
		helper.PanicIfError(err)

		programmingLanguage = progLangDomainRepo
	}

	questionDomain := domain.QuestionDomain{
		Id:                    questionResponse.Id,
		ProgrammingLanguageId: programmingLanguage.Id,
		Question:              questionReq.Question,
		CorrectAnswer:         questionReq.CorrectAnswer,
		UpdatedAt:             time.Now(),
	}

	incorrectAnswersDomain := []domain.IncorrectAnswerDomain{
		{QuestionId: questionResponse.Id, OptionA: questionReq.IncorrectOne},
		{QuestionId: questionResponse.Id, OptionB: questionReq.IncorrectTwo},
	}

	for _, incorrectAnswerDomain := range incorrectAnswersDomain {
		questionDomainResponse, err := service.backendRepository.UpdateQuestion(
			questionDomain, incorrectAnswerDomain,
		)
		helper.PanicIfError(err)

		question = questionDomainResponse
	}

	answersDomain, err := service.backendRepository.FindIncorrectAnswersByQuestionId(
		questionResponse.Id,
	)
	helper.PanicIfError(err)

	return helper.ToQuestionUpdateResponse(programmingLanguage, question, answersDomain), nil
}

func (service *BackendServiceImpl) DeleteQuestion(questionId int32) (bool, error) {
	_, err := service.backendRepository.DeleteQuestion(questionId)
	helper.PanicIfError(err)

	return true, nil
}
