package repositories

import (
	"sample/models"
	"sample/storage"
)

func InsertUser(user models.User) error {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO users (username,password) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(sqlStatement, user.Username, user.Password).Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}
func GetUserByUsername(username string) (*models.User, error) {
	db := storage.GetDB()
	user := models.User{}
	query := "SELECT * FROM users WHERE username=$1"
	err := db.QueryRow(query, username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func UpdatePassword(username, password string) error {
	db := storage.GetDB()
	query := "update users set password=$1 where username=$2"
	_, err := db.Exec(query, password, username)
	return err
}
