package model

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法")
	// t.Run("测试添加用户:", testAddUser)
	t.Run("测试获取用户:", testGetUserByID)
	t.Run("测试获取用户:", testGetUsers)
}

//如果函数名不是以Test开头，那么该函数默认不执行
func testAddUser(t *testing.T) {
	fmt.Println("子测试函数执行：")
	user := &User{}
	user.AddUser()
	user.AddUser2()
}

func testGetUserByID(t *testing.T) {
	fmt.Println("测试查询一条记录：")
	user := User{
		ID: 1,
	}
	u, _ := user.GetUserByID()
	fmt.Println("得到的User的信息是：", u)
}

func testGetUsers(t *testing.T) {
	fmt.Println("测试查询所有记录：")
	user := &User{}
	us, _ := user.GetUsers()
	for k, v := range us {
		fmt.Printf("第%v个用户是:%v\n", k+1, v)
	}

}