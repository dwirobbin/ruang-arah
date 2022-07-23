package web

import (
	"errors"
	"net"
	"net/mail"
	"strings"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type QuestionRequest struct {
	ProgrammingLanguage string `json:"programming_language"`
	Question            string `json:"question"`
	CorrectAnswer       string `json:"correct_answer"`
	IncorrectOne        string `json:"incorrect_one"`
	IncorrectTwo        string `json:"incorrect_two"`
}

type AnswerAttemptRequest struct {
	Answers []DataAttemptRequest `json:"answers"`
}

type DataAttemptRequest struct {
	QuestionId int32  `json:"question_id"`
	Answer     string `json:"answer"`
}

var (
	ErrRequiredUsername            = errors.New("required username")
	ErrRequiredEmail               = errors.New("required email")
	ErrRequiredPassword            = errors.New("required password")
	ErrInvalidEmail                = errors.New("invalid email")
	ErrDomainNotFound              = errors.New("domain not found")
	ErrRequiredQuestion            = errors.New("required Question")
	ErrRequiredProgrammingLanguage = errors.New("required Programming Language")
	ErrRequiredIncorrectOne        = errors.New("required Incorrect One")
	ErrRequiredIncorrectTwo        = errors.New("required Incorrect Two")
	ErrRequiredCorrectAnswer       = errors.New("required CorrectAnswer")
	ErrRequiredAnswer              = errors.New("required Answer")
	ErrInvalidAction               = errors.New("invalid action")
	ErrRequiredUpdate              = errors.New("programming language, question, correct answer, incorrect one, incorrect two is required")
	ErrRequiredQuestionId          = errors.New("required Question Id")
)

func (l *LoginRequest) ValidateLogin() error {
	if err := validateEmail(l.Email); err != nil {
		return err
	}
	if l.Password == "" {
		return ErrRequiredPassword
	}

	return nil
}

func (r *RegisterRequest) ValidateRegister() error {
	if r.Username == "" {
		return ErrRequiredUsername
	}
	if err := validateEmail(r.Email); err != nil {
		return err
	}
	if r.Password == "" {
		return ErrRequiredPassword
	}

	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return ErrRequiredEmail
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return ErrInvalidEmail
	}

	parts := strings.Split(email, "@")

	_, err = net.LookupMX(parts[1])
	if err != nil {
		return ErrDomainNotFound
	}

	return nil
}

func (q *QuestionRequest) ValidateQuestion(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if q.ProgrammingLanguage == "" {
			return ErrRequiredProgrammingLanguage
		}
		if q.Question == "" {
			return ErrRequiredQuestion
		}
		if q.CorrectAnswer == "" {
			return ErrRequiredCorrectAnswer
		}
		if q.IncorrectOne == "" {
			return ErrRequiredIncorrectOne
		}
		if q.IncorrectTwo == "" {
			return ErrRequiredIncorrectTwo
		}
		return nil
	case "update":
		if q.Question == "" && q.ProgrammingLanguage == "" && q.IncorrectOne == "" && q.IncorrectTwo == "" && q.CorrectAnswer == "" {
			return ErrRequiredUpdate
		}
		return nil
	}

	return ErrInvalidAction
}

func (a *AnswerAttemptRequest) ValidateAnswerAttempt() error {
	for _, v := range a.Answers {
		if err := v.ValidateDataAttempt(); err != nil {
			return err
		}
	}

	return nil
}

func (d *DataAttemptRequest) ValidateDataAttempt() error {
	if d.QuestionId == 0 {
		return ErrRequiredQuestion
	}
	if d.Answer == "" {
		return ErrRequiredAnswer
	}
	return nil
}
