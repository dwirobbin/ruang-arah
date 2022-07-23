package repository

import (
	"database/sql"
	"ruang-arah/backend/helper"
	"ruang-arah/backend/model/domain"
	"strings"
)

type BackendRepositoryImpl struct {
	DB *sql.DB
}

func NewBackendRepository(db *sql.DB) *BackendRepositoryImpl {
	return &BackendRepositoryImpl{
		DB: db,
	}
}

func (repo *BackendRepositoryImpl) FindProgrammingLanguageByName(name string) (domain.ProgrammingLanguageDomain, error) {
	var programmingLanguageDomain domain.ProgrammingLanguageDomain

	query := `SELECT id, name FROM programming_languages WHERE name = ?`
	row := repo.DB.QueryRow(query, name)
	err := row.Scan(&programmingLanguageDomain.Id, &programmingLanguageDomain.Name)
	helper.PanicIfError(err)

	return programmingLanguageDomain, nil
}

func (repo *BackendRepositoryImpl) FindProgrammingLanguageById(id int32) (domain.ProgrammingLanguageDomain, error) {
	var programmingLanguageDomain domain.ProgrammingLanguageDomain

	query := `SELECT id, name FROM programming_languages WHERE id = ?`
	row := repo.DB.QueryRow(query, id)
	err := row.Scan(&programmingLanguageDomain.Id, &programmingLanguageDomain.Name)
	helper.PanicIfError(err)

	return programmingLanguageDomain, nil
}

func (repo *BackendRepositoryImpl) FindProgrammingLanguageByQuestions(questions []domain.QuestionDomain) ([]domain.ProgrammingLanguageDomain, error) {
	var programmingLanguageDomain []domain.ProgrammingLanguageDomain

	query := `SELECT id, name FROM programming_languages WHERE id = ?;`

	for _, question := range questions {
		var programmingLanguage domain.ProgrammingLanguageDomain
		row := repo.DB.QueryRow(query, question.ProgrammingLanguageId)
		err := row.Scan(&programmingLanguage.Id, &programmingLanguage.Name)
		helper.PanicIfError(err)

		programmingLanguageDomain = append(programmingLanguageDomain, programmingLanguage)
	}

	return programmingLanguageDomain, nil
}

func (repo *BackendRepositoryImpl) SaveQuestion(questionDomain domain.QuestionDomain, incorrectAnswer domain.IncorrectAnswerDomain) (domain.QuestionDomain, error) {
	tx, err := repo.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := `
	INSERT INTO questions (proglang_id, question, correct_answer) VALUES (?, ?, ?);`

	result, err := tx.Exec(query,
		questionDomain.ProgrammingLanguageId, questionDomain.Question,
		questionDomain.CorrectAnswer,
	)
	helper.PanicIfError(err)

	questionId, err := result.LastInsertId()
	helper.PanicIfError(err)

	questionDomain.Id = int32(questionId)

	query = `
	INSERT INTO incorrect_answers (question_id, option_a, option_b) VALUES (?, ?, ?);`

	result, err = tx.Exec(
		query, questionId, incorrectAnswer.OptionA, incorrectAnswer.OptionB,
	)
	helper.PanicIfError(err)

	_, err = result.LastInsertId()
	helper.PanicIfError(err)

	return questionDomain, nil
}

func (repo *BackendRepositoryImpl) FindAllQuestions() ([]domain.QuestionDomain, error) {
	var questions []domain.QuestionDomain

	query := `SELECT id, proglang_id, question FROM questions;`

	rows, err := repo.DB.Query(query)
	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		var question domain.QuestionDomain
		err := rows.Scan(&question.Id, &question.ProgrammingLanguageId, &question.Question)
		helper.PanicIfError(err)

		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}

func (repo *BackendRepositoryImpl) FindQuestionById(id int32) (domain.QuestionDomain, error) {
	var question domain.QuestionDomain

	query := `SELECT id, proglang_id, question, correct_answer FROM questions WHERE id = ?`

	row := repo.DB.QueryRow(query, id)
	err := row.Scan(
		&question.Id, &question.ProgrammingLanguageId,
		&question.Question, &question.CorrectAnswer,
	)
	helper.PanicIfError(err)

	return question, nil
}

func (repo *BackendRepositoryImpl) FindIncorrectAnswersByQuestionId(questionId int32) (domain.IncorrectAnswerDomain, error) {
	query := `
	SELECT id, question_id, option_a, option_b 
	FROM incorrect_answers WHERE question_id = ?;`

	var incorrectAnswerDomain domain.IncorrectAnswerDomain
	row := repo.DB.QueryRow(query, questionId)
	err := row.Scan(
		&incorrectAnswerDomain.Id, &incorrectAnswerDomain.QuestionId,
		&incorrectAnswerDomain.OptionA, &incorrectAnswerDomain.OptionB,
	)
	helper.PanicIfError(err)

	return incorrectAnswerDomain, nil
}

func (repo *BackendRepositoryImpl) UpdateQuestion(questionDomain domain.QuestionDomain, incorrectAnswerDomain domain.IncorrectAnswerDomain) (domain.QuestionDomain, error) {
	tx, err := repo.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// begin update question
	query := `UPDATE questions SET `

	qParts := []string{}
	args := []interface{}{}

	if questionDomain.Question != "" {
		qParts = append(qParts, "question = ?")
		args = append(args, questionDomain.Question)
	}
	if questionDomain.CorrectAnswer != "" {
		qParts = append(qParts, "correct_answer = ?")
		args = append(args, questionDomain.CorrectAnswer)
	}
	if questionDomain.ProgrammingLanguageId != 0 {
		qParts = append(qParts, "proglang_id = ?")
		args = append(args, questionDomain.ProgrammingLanguageId)
	}

	query += strings.Join(qParts, ", ")
	query += ", updated_at = ? WHERE id = ?;"

	args = append(args, questionDomain.UpdatedAt, questionDomain.Id)

	_, err = tx.Exec(query, args...)
	helper.PanicIfError(err)
	// end update question

	// begin update answer
	query = `UPDATE incorrect_answers SET `

	qParts = []string{}
	args = []interface{}{}

	if incorrectAnswerDomain.OptionA != "" {
		qParts = append(qParts, "option_a = ?")
		args = append(args, incorrectAnswerDomain.OptionA)
	}
	if incorrectAnswerDomain.OptionB != "" {
		qParts = append(qParts, "option_b = ?")
		args = append(args, incorrectAnswerDomain.OptionB)
	}

	query += strings.Join(qParts, ", ")
	query += ", updated_at = ? WHERE question_id = ?;"

	args = append(args, incorrectAnswerDomain.UpdatedAt, questionDomain.Id)

	_, err = tx.Exec(query, args...)
	helper.PanicIfError(err)
	// end update answer

	return questionDomain, nil
}

func (repo *BackendRepositoryImpl) DeleteQuestion(questionId int32) (bool, error) {
	query := `DELETE FROM questions WHERE id = ?;`

	_, err := repo.DB.Exec(query, questionId)
	helper.PanicIfError(err)

	query = `DELETE FROM incorrect_answers WHERE question_id = ?;`
	_, err = repo.DB.Exec(query, questionId)
	helper.PanicIfError(err)

	return true, nil
}
