package models

import (
	"jishiben/dao"
	"golang.org/x/crypto/bcrypt"
)

// Todo Model
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}
type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Nickname       string
	PasswordDigest string
}
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}
type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}
type UserResponse struct {
	Response
	Data User `json:"data"`
}
const (PassWordCost = 12)
	// PassWordCost 密码加密难度

func BuildUser(user *User) User {
	return User{
		ID:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,

	}
}
// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

// BuildListResponse 列表构建器
func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Data: DataList{
			Items: items,
			Total: total,
		},
	}
}
func BuildUserResponse(user *User) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}

/*
	Todo这个Model的增删改查操作都放在这里
 */
// CreateATodo 创建todo
func CreateATodo(todo *Todo) (err error){
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todoList []*Todo, err error){
	if err = dao.DB.Find(&todoList).Error; err != nil{
		return nil, err
	}
	return
}

func GetATodo(id string)(todo *Todo, err error){
	todo = new(Todo)
	if err = dao.DB.Debug().Where("id=?", id).First(todo).Error; err!=nil{
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo)(err error){
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string)(err error){
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}

