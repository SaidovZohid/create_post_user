package main

import (
	"database/sql"
	"fmt"
	"time"
)

type DbManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) *DbManager {
	return &DbManager{db: db}
}

type User struct {
	Id int
	FirstName string
	SecondName string
	Age uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Post struct {
	Id int
	UserId int
	UserPost string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
} 

type GetUserStruct struct {
	Id int
	FirstName string
	SecondName string
	Age uint8
	posts []Post
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

func (d *DbManager) GetUser(id int) (*GetUserStruct, error){
	// var user_posts *GetUserStruct
	query := `
		SELECT 
			users.id,
			first_name,
			second_name,
			age,
			user_post
		FROM users JOIN posts ON users.id = user_id;
	`
	row , err:= d.db.Query(query)
	fmt.Println(row)
	if err != nil {
		return nil, err
	}
	return nil, err
}