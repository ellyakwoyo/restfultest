package services

import (
	"restfultest/database"
	"restfultest/models"
)

func GetAllUsers() ([]models.User, error) {
	query := "SELECT id, first_name, last_name, email, gender, ip_address FROM users"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Gender, &user.IpAddress)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(user *models.User) error {
	query := "INSERT INTO users (first_name, last_name, email, gender, ip_address, password) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := database.DB.Exec(query, user.FirstName, user.LastName, user.Email, user.Gender, user.IpAddress, user.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}
