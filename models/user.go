package models

import (
	"database/sql"
	"fmt"
	"ginex/database"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id            int64
	Name          string
	Email         string
	Password      string
	RememberToken *string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

func Get() {
	user := User{}
	/*db := database.Init()
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", 1)
	err := row.Scan(&user.Id, &user.Name, &user.Email,&user.Password, &user.RememberToken, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		fmt.Println(err)
	}*/
	fmt.Println(user)
}

func GetUser(email string) (User,error) {
	user := User{}
	db := database.Init()
	row := db.QueryRow("SELECT * FROM users WHERE email = ? LIMIT 1", email)
	err := row.Scan(&user.Id, &user.Name, &user.Email,&user.Password, &user.RememberToken, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	return user,err
}

func StoreUser(user User) (int64, error) {
	//初始化数据库
	db := database.Init()

	//获取这个邮箱是否存在
	_,err := GetUser(user.Email)
	fmt.Println(err)

	//如果错误不是没找到数据（那就是找到了数据）就返回不继续执行
	if err != sql.ErrNoRows {
		return 0,err
	}

	//插入数据
	stmt,_ := db.Prepare("INSERT INTO users(`name`,`email`,`password`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)")
	now := time.Now().Format("2006-01-02 15:04:05")
	passwordHashed,_ := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)

	result,_ := stmt.Exec(user.Name,user.Email,passwordHashed,now,now)
	id,err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return id,err
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
