package main

import (
	"database/sql"
	"time"
)

type DbManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) *DbManager {
	return &DbManager{db: db}
}

type User struct {
	Id         int
	FirstName  string
	SecondName string
	Age        uint8
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type Post struct {
	Id        int
	UserId    int
	UserPost  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type GetUserStruct struct {
	Id         int
	FirstName  string
	SecondName string
	Age        uint8
	posts      string
}

func (d *DbManager) CreateUser(u *User) (*User, error) {
	var user User
	query := `
		INSERT INTO users(
			first_name,
			second_name,
			age
		) VALUES ($1, $2, $3)
		RETURNING id, first_name, second_name, age, created_at
	`
	row := d.db.QueryRow(
		query,
		u.FirstName,
		u.SecondName,
		u.Age,
	)
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.SecondName,
		&user.Age,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *DbManager) CreatePost(p *Post) (*Post, error){
	var post Post
	query := `
		INSERT INTO posts(
			user_id, 
			user_post
		) VALUES (
			$1,
			$2
		) RETURNING id,user_id,user_post, created_at
	`
	row := d.db.QueryRow(
		query,
		p.UserId,
		p.UserPost,
	)
	err := row.Scan(
		&post.Id,
		&post.UserId,
		&post.UserPost,
		&post.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (d *DbManager) GetUser(id int) ([]*GetUserStruct, error) {
	var user_posts []*GetUserStruct
	query := `
		SELECT 
			users.id,
			first_name,
			second_name,
			age,
			user_post
		FROM users JOIN posts ON users.id = user_id;
	`
	row, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var post GetUserStruct
		err = row.Scan(
			&post.Id,
			&post.FirstName,
			&post.SecondName,
			&post.Age,
			&post.posts,
		)
		if err != nil {
			return nil, err
		}
		user_posts = append(user_posts, &post)
	}
	return user_posts, nil
}

func (d *DbManager) UpdateUser(u *User) (*User, error) {
	query := `
		UPDATE users SET 
			first_name = $1,
			second_name = $2,
			age = $3,
			updated_at = $4
		WHERE id = $5
		RETURNING id, first_name, second_name, age, updated_at
	`
	row := d.db.QueryRow(
		query,
		u.FirstName,
		u.SecondName,
		u.Age,
		u.UpdatedAt,
		u.Id,
	)
	var result User
	err := row.Scan(
		&result.Id,
		&result.FirstName,
		&result.SecondName,
		&result.Age,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
} 

func (d *DbManager) UpdatePost(p *Post) (*Post, error) {
	query := `
		UPDATE posts SET 
			user_post = $1,
			updated_at = $2 
		WHERE id = $3
		RETURNING id,user_id,user_post, updated_at
	`
	row := d.db.QueryRow(
		query,
		p.UserPost,
		p.UpdatedAt,
		p.Id,
	)
	var result Post
	err := row.Scan(
		&result.Id,
		&result.UserId,
		&result.UserPost,
		&result.UpdatedAt,
	)
	if err != nil  {
		return nil, err
	}
	return &result, err
}

func (d *DbManager) DeleteUser(u *User) {
	query := `
		UPDATE users SET 
			deleted_at = $1
		WHERE id = $2
	`
	d.db.QueryRow(
		query,
		u.DeletedAt,
		u.Id,
	)
}

func (d *DbManager) DeletedPost(p *Post) {
	query := `
		UPDATE posts SET 
			deleted_at = $1
		WHERE id = $2 and user_id = $3
	`
	d.db.QueryRow(
		query,
		p.DeletedAt,
		p.Id,
		p.UserId,
	)
}