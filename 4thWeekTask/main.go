package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 用户
type User struct {
	Username string
	Password string
	Nickname string
}

// 存储信息
var users = make(map[string]User)

// 存储登录状态：token -> username
var sessions = make(map[string]string)

// 注册
func register(w http.ResponseWriter, r *http.Request) {
	var requser User
	json.NewDecoder(r.Body).Decode(&requser)
	if _, ok := users[requser.Username]; ok {
		fmt.Fprintln(w, "用户名已存在")
		return
	}
	users[requser.Username] = requser
	fmt.Fprintln(w, "注册成功")
}

// 登录
func login(w http.ResponseWriter, r *http.Request) {
	var requser User
	json.NewDecoder(r.Body).Decode(&requser)
	u, ok := users[requser.Username]
	if !ok || u.Password != requser.Password {
		fmt.Fprintln(w, "用户名或密码错误")
		return
	}
	token := requser.Username + "123456"
	sessions[token] = requser.Username
	fmt.Fprintln(w, token)
}

// 获取信息
func getInfo(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	username := sessions[token]
	user := users[username]
	user.Password = ""
	json.NewEncoder(w).Encode(user)
}

// 修改信息
func updateInfo(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	username := sessions[token]
	var data struct {
		Nickname string
		Age      int
	}
	json.NewDecoder(r.Body).Decode(&data)
	user := users[username]
	user.Nickname = data.Nickname
	users[username] = user
	fmt.Fprintln(w, "修改成功")
}

// 修改密码
func updatePassword(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	username := sessions[token]
	var data struct {
		OldPassword string
		NewPassword string
	}
	json.NewDecoder(r.Body).Decode(&data)
	user := users[username]
	if user.Password != data.OldPassword {
		fmt.Fprintln(w, "原密码错误")
		return
	}
	user.Password = data.NewPassword
	users[username] = user
	fmt.Fprintln(w, "密码修改成功")
} 
func main() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/info", getInfo)
	http.HandleFunc("/updateInfo", updateInfo)
	http.HandleFunc("/updatePassword", updatePassword)
	http.ListenAndServe(":8080", nil)
}

//稍微借助了ai进行学习与指导
