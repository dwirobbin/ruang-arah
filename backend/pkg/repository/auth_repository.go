package repository

import (
	"database/sql"
	"errors"
	"ruang-arah/backend/helper"
	"ruang-arah/backend/model/domain"
)

type AuthRepositoryImpl struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		DB: db,
	}
}

func (repo *AuthRepositoryImpl) FindUser(email, password string) (domain.UserDomain, error) {
	query := `
	SELECT id, username, email, password, role, loggedin
	FROM users 
	WHERE email = ? AND password = ?;`

	var user domain.UserDomain
	row := repo.DB.QueryRow(query, email, password)
	err := row.Scan(
		&user.Id, &user.Username, &user.Email,
		&user.Password, &user.Role, &user.Loggedin,
	)
	if err != nil {
		return user, errors.New("login failed")
	}

	query = `UPDATE users SET loggedin = true WHERE id = ?;`
	_, err = repo.DB.Exec(query, user.Id)
	helper.PanicIfError(err)

	return user, nil
}

func (repo *AuthRepositoryImpl) Save(userDomain domain.UserDomain, email string) (domain.UserDomain, error) {
	query := `SELECT id, email FROM users WHERE email = ?`
	row := repo.DB.QueryRow(query, email)
	err := row.Scan(&userDomain.Id, &userDomain.Email)
	if err == nil {
		return userDomain, errors.New("email already exists")
	}

	query = `
	INSERT INTO users 
	(username, email, password, role, loggedin, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?);`

	result, err := repo.DB.Exec(query,
		userDomain.Username, email, userDomain.Password, userDomain.Role,
		userDomain.Loggedin, userDomain.CreatedAt, userDomain.UpdatedAt,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	userDomain.Id = int32(id)
	return userDomain, nil
}

func (repo AuthRepositoryImpl) FindUserById(id int32) (domain.UserDomain, error) {
	query := `SELECT id, username, email, role FROM users WHERE id = ?;`

	var user domain.UserDomain
	row := repo.DB.QueryRow(query, id)
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Role)
	if err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}
