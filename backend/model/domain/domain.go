package domain

import "time"

type UserDomain struct {
	Id        int32
	Username  string
	Email     string
	Password  string
	Role      string
	Loggedin  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProgrammingLanguageDomain struct {
	Id        int32
	Name      string
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type QuestionDomain struct {
	Id                    int32
	ProgrammingLanguageId int32
	Question              string
	CorrectAnswer         string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type IncorrectAnswerDomain struct {
	Id         int32
	QuestionId int32
	OptionA    string
	OptionB    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type AnswerAttemptDomain struct {
	Id         int32
	Answer     string
	QuestionId int32
	UserId     int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type LevelDomain struct {
	Id        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RecommendationDomain struct {
	Id        int32
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
