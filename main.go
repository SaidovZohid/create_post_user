package main

import (
	"database/sql"
	"fmt"
	// "log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "post"
)

func main(){
	connstr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connstr)
	if err != nil {
	  panic(err)
	}
	fmt.Println("Database Connected")
	defer db.Close()
	err = db.Ping()
	if err != nil {
	  panic(err)
	}
	dbManager := NewDBManager(db)
	// _, err = dbManager.CreateUser(&User{
	// 	FirstName: "Zohid",
	// 	SecondName: "Saidov",
	// 	Age: 18,
	// })
	// if err != nil {
	// 	log.Fatalf("Error while creating new user: %v", err)
	// }
	// post, err := dbManager.CreatePost(&Post{
	// 	UserId: 1,
	// 	UserPost: "I did it this job",
	// })
	// if err != nil {
	// 	log.Fatalf("Error while creating new post: %v", err)
	// }
	// fmt.Println(post)
	// user_post, err := dbManager.GetUser(1)
	// if err != nil {
	// 	log.Fatalf("Error while getting user and posts: %v", err)
	// }
	// for _, v := range user_post {
	// 	fmt.Println(*v)
	// }
	// user, err := dbManager.UpdateUser(&User{
	// 	Id: 1,
	// 	FirstName: "Ismoiljon",
	// 	SecondName: "Rustamov",
	// 	Age: 16,
	// 	UpdatedAt: time.Now(),
	// })
	// if err != nil {
	// 	log.Fatalf("Error while updating user: %v\n", err)
	// }
	// fmt.Println(*user)
	// post, err := dbManager.UpdatePost(&Post{
	// 	Id: 3,
	// 	UserId: 1,
	// 	UserPost: "Hp Pavilion Gaming",
	// 	UpdatedAt: time.Now(),
	// })
	// if err != nil {
	// 	log.Fatalf("Error while updating post: %v\n", err)
	// }
	// fmt.Println(*post)
	// dbManager.DeleteUser(&User{
	// 	Id: 1,
	// 	DeletedAt: time.Now(),
	// })
	dbManager.DeletedPost(&Post{
		Id: 4,
		UserId: 1,
		DeletedAt: time.Now(),
	})
}

