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

	sql := "SELECT id, email, password FROM users WHERE email = $1"
	err = db.QueryRow(sql, email).Scan(&usr.ID, &usr.Email, &usr.Password)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}
