package helper

import (
	"ruang-arah/backend/model/domain"
	"ruang-arah/backend/model/web"
)

func ToLoginResponse(userDomain domain.UserDomain, token string) web.LoginResponse {
	return web.LoginResponse{
		Id:       userDomain.Id,
		Username: userDomain.Username,
		Role:     userDomain.Role,
		Token:    token,
	}
}

func ToRegisterResponse(userDomain domain.UserDomain, email string) web.RegisterResponse {
	return web.RegisterResponse{
		Username: userDomain.Username,
		Email:    email,
		Role:     userDomain.Role,
	}
}

func ToQuestionResponse(proglangDomain domain.ProgrammingLanguageDomain, questionDomain domain.QuestionDomain, incorrectAnswer domain.IncorrectAnswerDomain) web.QuestionResponse {
	return web.QuestionResponse{
		Id:                  questionDomain.Id,
		ProgrammingLanguage: proglangDomain.Name,
		Question:            questionDomain.Question,
		CorrectAnswer:       questionDomain.CorrectAnswer,
		IncorrectAnswers:    FromAnswer(incorrectAnswer),
	}
}

func ToQuestionResponses(questions []domain.QuestionDomain, incorrectAnswer domain.IncorrectAnswerDomain, programingLanguage domain.ProgrammingLanguageDomain) []web.QuestionResponse {
	var questionListResponse []web.QuestionResponse
	for _, question := range questions {
		questionListResponse = append(questionListResponse, web.QuestionResponse{
			Id:                  question.Id,
			ProgrammingLanguage: programingLanguage.Name,
			Question:            question.Question,
			CorrectAnswer:       question.CorrectAnswer,
			IncorrectAnswers:    FromAnswer(incorrectAnswer),
		})
	}
	return questionListResponse
}

func FromAnswer(incorrectAnswer domain.IncorrectAnswerDomain) []string {
	var incorrectAnswerResponses []string
	incorrectAnswerResponses = append(incorrectAnswerResponses,
		incorrectAnswer.OptionA, incorrectAnswer.OptionB,
	)

	return incorrectAnswerResponses
}

func ToQuestionListResponses(questions []domain.QuestionDomain, programingLanguages []domain.ProgrammingLanguageDomain) []web.QuestionListResponse {
	var questionListResponse []web.QuestionListResponse

	for _, question := range questions {
		var programmingLanguageName string
		for _, progLang := range programingLanguages {
			if question.ProgrammingLanguageId == progLang.Id {
				programmingLanguageName = progLang.Name
			}
		}

		questionListResponse = append(questionListResponse, web.QuestionListResponse{
			Id:                  question.Id,
			Question:            question.Question,
			ProgrammingLanguage: programmingLanguageName,
		})
	}

	return questionListResponse
}

func ToQuestionUpdateResponse(proglangDomain domain.ProgrammingLanguageDomain, questionDomain domain.QuestionDomain, incorrectAnswer domain.IncorrectAnswerDomain) web.QuestionResponse {
	return web.QuestionResponse{
		Id:                  questionDomain.Id,
		ProgrammingLanguage: proglangDomain.Name,
		Question:            questionDomain.Question,
		CorrectAnswer:       questionDomain.CorrectAnswer,
		IncorrectAnswers:    FromAnswer(incorrectAnswer),
	}
}

func ToTotalQuestionsResponse(totalQuestions int32) web.TotalQuestionsResponse {
	return web.TotalQuestionsResponse{
		TotalQuestions: totalQuestions,
	}
}

func ToProgrammingLanguageResponses(programmingLanguages []domain.ProgrammingLanguageDomain) []web.ProgrammingLanguageResponse {
	var programmingLanguageResponses []web.ProgrammingLanguageResponse
	for _, programmingLanguage := range programmingLanguages {
		programmingLanguageResponses = append(programmingLanguageResponses, web.ProgrammingLanguageResponse{
			Id:       programmingLanguage.Id,
			Name:     programmingLanguage.Name,
			ImageUrl: programmingLanguage.ImageUrl,
		})
	}

	return programmingLanguageResponses
}

func ToAnswersAttemptsResponse(user domain.UserDomain, level domain.LevelDomain, score int32) web.AnswersAttemptsResponse {
	return web.AnswersAttemptsResponse{
		LevelId:  level.Id,
		Username: user.Username,
		Level:    level.Name,
		Score:    score,
	}
}

func ToRecommendationResponse(recommendation domain.RecommendationDomain) web.RecommendationResponse {
	return web.RecommendationResponse{
		ImageUrl: recommendation.ImageUrl,
	}
}
