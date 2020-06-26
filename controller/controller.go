package controller

import (
	"jishiben/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"jishiben/api"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func UserRegist(c *gin.Context){
	var server api.UserRegister
	if err:=c.ShouldBind(&server); err == nil {
	if user,err:=server.Register();err!=nil{
	c.JSON(200,err)
	}else{
	res:=models.BuildUserResponse(&user)
	c.JSON(200,res)
	}

	}else{
		c.JSON(200,nil)
	}
	c.HTML(http.StatusOK, "login.html", nil)
	}
func UserLogin(c *gin.Context){
	var service api.UserLog
	if err:=c.ShouldBind(&service);err!=nil{
		if user,err:=service.Login();err!=nil{
			c.JSON(200,err)
	}else{
		res := models.BuildUserResponse(&user)
		c.JSON(200, res)
	}
}else{
		c.JSON(200, nil)
}
	c.HTML(http.StatusOK, "regist.html", nil)
}
func CreateTodo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	// 2. 存入数据库
	err:=models.CreateATodo(&todo)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 2000,
		//	"msg": "success",
		//	"data": todo,
		//})
	}
}

func GetTodoList(c *gin.Context) {
	// 查询todo这个表里的所有数据
	todoList, err := models.GetAllTodo()
	if err!= nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, todoList)
	}
}
//修改代办事项
func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err!= nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}
//删除
func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteATodo(id);err!=nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, gin.H{id:"deleted"})
	}
}