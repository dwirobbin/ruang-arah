package web

type LoginResponse struct {
	Id       int32  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type QuestionResponse struct {
	Id                  int32    `json:"id"`
	ProgrammingLanguage string   `json:"programming_language"`
	Question            string   `json:"question"`
	CorrectAnswer       string   `json:"correct_answer"`
	IncorrectAnswers    []string `json:"incorrect_answers"`
}

type QuestionListResponse struct {
	Id                  int32  `json:"id"`
	Question            string `json:"question"`
	ProgrammingLanguage string `json:"programming_language"`
}

type TotalQuestionsResponse struct {
	TotalQuestions int32 `json:"total_questions"`
}

type ProgrammingLanguageResponse struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type AnswersAttemptsResponse struct {
	LevelId  int32  `json:"level_id"`
	Username string `json:"username"`
	Level    string `json:"level"`
	Score    int32  `json:"score"`
}

type RecommendationResponse struct {
	ImageUrl string `json:"image_url"`
}
