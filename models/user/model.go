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
