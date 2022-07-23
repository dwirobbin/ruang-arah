package repository

import (
	"database/sql"
	"ruang-arah/backend/helper"
	"ruang-arah/backend/model/domain"
)

type FrontendRepositoryImpl struct {
	DB *sql.DB
}

func NewFrontendRepository(db *sql.DB) *FrontendRepositoryImpl {
	return &FrontendRepositoryImpl{
		DB: db,
	}
}

func (repo *FrontendRepositoryImpl) FindAllProgrammingLanguages() ([]domain.ProgrammingLanguageDomain, error) {
	var programmingLanguages []domain.ProgrammingLanguageDomain

	query := `SELECT id, name, image_url FROM programming_languages`
	rows, err := repo.DB.Query(query)
	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		var programmingLanguageDomain domain.ProgrammingLanguageDomain
		err := rows.Scan(
			&programmingLanguageDomain.Id, &programmingLanguageDomain.Name,
			&programmingLanguageDomain.ImageUrl,
		)
		helper.PanicIfError(err)

		programmingLanguages = append(programmingLanguages, programmingLanguageDomain)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return programmingLanguages, nil
}

func (repo *FrontendRepositoryImpl) FindAllQuestionByProgrammingLanguageIdWithPagination(programmingLanguageId, page, limit int32) ([]domain.QuestionDomain, error) {
	var questions []domain.QuestionDomain

	query := `
	SELECT id, proglang_id, question, correct_answer FROM questions 
	WHERE proglang_id = ? ORDER BY id LIMIT ? OFFSET ?;`

	rows, err := repo.DB.Query(query, programmingLanguageId, limit, (page-1)*limit)
	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		var questionDomain domain.QuestionDomain
		err := rows.Scan(
			&questionDomain.Id, &questionDomain.ProgrammingLanguageId,
			&questionDomain.Question, &questionDomain.CorrectAnswer,
		)
		helper.PanicIfError(err)

		questions = append(questions, questionDomain)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return questions, nil
}

func (repo *FrontendRepositoryImpl) SaveAnswersAttempts(userId int32, answersAttempts []domain.AnswerAttemptDomain) (bool, error) {
	query := `DELETE FROM answers_attempts;`
	_, err := repo.DB.Exec(query)
	helper.PanicIfError(err)

	query = `INSERT INTO answers_attempts (answer, question_id, user_id) VALUES (?, ?, ?)`

	for _, answerAttempt := range answersAttempts {
		_, err := repo.DB.Exec(query, answerAttempt.Answer, answerAttempt.QuestionId, userId)
		helper.PanicIfError(err)
	}

	return true, nil
}

func (repo *FrontendRepositoryImpl) FindTotalAnswersAttemptsByUserId(userId int32) (int32, error) {
	var total int32

	query := `
	SELECT COUNT(aa.question_id) AS "correct_amount" FROM answers_attempts AS aa
	INNER JOIN questions AS q ON aa.question_id = q.id
	WHERE q.correct_answer = aa.answer AND aa.user_id = ?;`

	row := repo.DB.QueryRow(query, userId)
	err := row.Scan(&total)
	helper.PanicIfError(err)

	return total * 10, nil
}

func (repo *FrontendRepositoryImpl) FindLevelByName(name string) (domain.LevelDomain, error) {
	var level domain.LevelDomain

	query := `SELECT id, name FROM levels WHERE name = ?`

	row := repo.DB.QueryRow(query, name)
	err := row.Scan(&level.Id, &level.Name)
	helper.PanicIfError(err)

	return level, nil
}

func (repo *FrontendRepositoryImpl) FindRecommendationByLevelId(levelId int32) (domain.RecommendationDomain, error) {
	var recommendation domain.RecommendationDomain

	query := `SELECT id, image_url FROM recommendations WHERE level_id = ?`

	row := repo.DB.QueryRow(query, levelId)
	err := row.Scan(&recommendation.Id, &recommendation.ImageUrl)
	helper.PanicIfError(err)

	return recommendation, nil
}
