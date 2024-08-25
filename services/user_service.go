package services

import (
	"errors"
	"restfultest/database"
	"restfultest/models"
	"time"
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

func GetUser(id int) (*models.User, error) {
	// Define the query to retrieve a user by ID
	query := "SELECT id, first_name, last_name, email, gender, ip_address FROM users WHERE id = ?"

	// Execute the query using QueryRow since we expect a single result
	row := database.DB.QueryRow(query, id)

	// Create a User object to hold the result
	var user models.User

	// Scan the result into the User struct
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Gender, &user.IpAddress)
	if err != nil {
		return nil, err
	}

	// Return the retrieved user
	return &user, nil
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

func UpdateUser(user *models.User) error {
	currentTime := time.Now()
	query := `UPDATE users 
	          SET first_name = ?, last_name = ?, email = ?, gender = ?, ip_address = ?, password = ?, updated_at = ?
	          WHERE id = ?`

	// Execute the update query
	_, err := database.DB.Exec(query, user.FirstName, user.LastName, user.Email, user.Gender, user.IpAddress, user.Password, currentTime, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {

	queryCheck := "SELECT COUNT(*) FROM users WHERE id = ?"
	var count int
	err := database.DB.QueryRow(queryCheck, id).Scan(&count)
	if err != nil {
		return errors.New("Fsiled to ckeck if user exists")
	}

	if count == 0 {
		return errors.New("User not found")
	}

	// Step 2: Delete the user
	queryDelete := "DELETE FROM users WHERE id = ?"
	_, err = database.DB.Exec(queryDelete, id)
	if err != nil {
		return errors.New("Failed to delete user")
	}

	return nil
}
