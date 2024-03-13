package repository

import (
	"asset-tracker/model"
	"database/sql"
)

func AddUser(db *sql.DB, user model.User) (err error) {
	sql := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1 ,$2, $3, NOW(), NOW())"
	errs := db.QueryRow(sql, user.Name, user.Email, user.Password)

	return errs.Err()
}

func FindUserByEmail(db *sql.DB, email string) (user *model.User, err error) {
	var usr model.User

	sql := "SELECT id, email, name, password FROM users WHERE email = $1"
	err = db.QueryRow(sql, email).Scan(&usr.ID, &usr.Email, &usr.Name, &usr.Password)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func FindUserById(db *sql.DB, id int) (user *model.User, err error) {
	var usr model.User

	sql := "SELECT id, email, name, password FROM users WHERE id = $1"
	err = db.QueryRow(sql, id).Scan(&usr.ID, &usr.Email, &usr.Name, &usr.Password)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func UpdateUser(db *sql.DB, user model.UpdateProfileRequest) (err error) {
	var sql string

	if user.Password != nil {
		sql = "UPDATE users SET name = $2, email = $3, password = $4 WHERE id = $1"
		errs := db.QueryRow(sql, user.ID, user.Name, user.Email, user.Password)
		err = errs.Err()
	} else {
		sql = "UPDATE users SET name = $2, email = $3 WHERE id = $1"
		errs := db.QueryRow(sql, user.ID, user.Name, user.Email)
		err = errs.Err()
	}

	return
}
