package services

import (
	"restfultest/database"
	"restfultest/models"
)

func GetAllPosts() ([]models.Post, error) {
	query := "SELECT id, title, content FROM posts"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func CreatePost(post *models.Post) error {
	query := "INSERT INTO posts (title, content) VALUES (?, ?)"
	result, err := database.DB.Exec(query, post.Title, post.Content)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	post.ID = int(id)
	return nil
}
