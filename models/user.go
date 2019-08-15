package models

import (
	"time"
	"ginex/database"
	"fmt"
	"database/sql"
)

type User struct {
	Id            int
	Name          string
	Email         string
	Password      string
	RememberToken string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

var db *sql.DB

func init() {
	db = database.Init()
}

func Get() {
	user := User{}
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", 1)
	err := row.Scan(&user.Id, &user.Name, &user.Email,&user.Password, &user.RememberToken, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}

func GetUser(email string) {
	user := User{}
	row := db.QueryRow("SELECT * FROM users WHERE email = ?", email)
	err := row.Scan(&user.Id, &user.Name, &user.Email,&user.Password, &user.RememberToken, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}

func StoreUser(user User) (int64, error) {
	//todo store user data
	//return db.Exec("INSERT INTO users(`name`,`email`,`password`) VALUES (?,?,?)",user.Name,user.Email,user.Password)
}

/*func All() []User {
	var users []User
	rows,err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next(){
		var user User
		rows.Scan(&user.Id, &user.Name, &user.Email,&user.Password, &user.RememberToken, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		users = append(users,user)
	}
	return users
}*/
