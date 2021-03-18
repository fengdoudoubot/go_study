package model

import (
	"fmt"
	"go_code/project09/db"
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

func (user *User) AddUser() error {
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	inStmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常：", err)
		return err
	}
	_, err2 := inStmt.Exec("admin", "123456", "admin@qq.com")
	if err2 != nil {
		fmt.Println("执行出现异常：", err2)
		return err
	}
	return nil
}

func (user *User) AddUser2() error {
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	_, err := db.Db.Exec(sqlStr, "admin2", "666666", "admin2@sina.com")
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

func (user *User) GetUserByID() (*User, error) {
	sqlStr := "select id,username,password,email from users where id = ?"
	row := db.Db.QueryRow(sqlStr, user.ID)
	var id int
	var username string
	var password string
	var email string
	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		return nil, err
	}
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u, nil
}

func (user *User) GetUsers() ([]*User, error) {
	sqlStr := "select id,username,password,email from users"
	rows, err := db.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var users []*User
	for rows.Next() {
		var id int
		var username string
		var password string
		var email string
		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			return nil, err
		}
		u := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		users = append(users, u)
	}
	return users, nil
}
