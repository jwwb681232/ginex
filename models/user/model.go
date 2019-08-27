/*
package user

import (
	"ginex/database"
	"time"
	"strings"
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

func GetOne(c map[string]interface{},f []string) (User,error) {
	field := strings.Join(f,",")
	condition := ""
	for key,value := range c {
		condition += key + "=" + value.(string) + " AND "
	}

	user := User{}
	row := database.DB.QueryRow("SELECT " + field + " FROM users WHERE " + condition)
	err := row.Scan(&user.Id, &user.Name, &user.Email,&user.Password, &user.RememberToken, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return user,err
	}
	return user,err
}
*/
package user

import (
	"github.com/jinzhu/gorm"
	"ginex/database"
)

type User struct {
	gorm.Model
	Name          string  `gorm:"column:name"`
	Email         string  `gorm:"column:email"`
	Password      string  `gorm:"column:password"`
	RememberToken *string `gorm:"column:remember_token"`
}

func (User) TableName() string {
	return "users"
}

/*func (User) WhereEmail(email *string) (User,bool) {
	var userData User
	if err := database.Db.Where("email = ?", email).First(&userData).RecordNotFound();err != false {
		return userData,err
	}
	return userData,false
}*/

func (User) WhereEmail(email *string) (User, bool) {
	var userData User
	var notFound bool

	notFound = database.Db.Where("email = ?", email).First(&userData).RecordNotFound()
	return userData, notFound
}

func (User) CreateUser(u User) *gorm.DB {
	return database.Db.Create(&u)
}
