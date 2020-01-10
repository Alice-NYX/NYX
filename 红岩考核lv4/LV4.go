package main

import (
	"awesomeProject12/model"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var db *sql.DB
var Administrator string
var id string
var point int
func init() {
	db, _ = sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
}
func DBConn() *sql.DB{
	return  db
}

func main(){
	r:=gin.Default()
	r.POST("/register",Register)
	r.POST("/login",Login)
	r.POST("/insert",Insert)
	r.POST("/set",Set)
	r.Run(":8080")
}

func Register(c *gin.Context){
	username:=c.PostForm("username")
	password:=c.PostForm("password")

	fmt.Println("user:"+username+password)
	if model.UserSignup(username,password){
		c.JSON(500,gin.H{"status":http.StatusInternalServerError,"message":"数据库Insert报错"})
	}else {
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "注册成功"})
	}
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if model.UserSignin(username, password) {
		c.SetCookie("username", username, 10, "localhost:8080", "localhost", false, true)
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "登录成功"})
	} else {
		c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "登录失败，用户名或密码错误"})
	}
}

func UserSignup(username string,password string)bool{
	stmt,err:=DBConn().Prepare(
		"insert into user(username,password)values(?,?) ")
	if err!=nil{
		fmt.Println("fail to insert")
		return false
	}
	defer stmt.Close()

	_,err=stmt.Exec(username,password)
	if err!=nil{
		fmt.Println("fail to insert")
		return false
	}

	return false
}

func UserSignin( username string,password string)bool{
	stmt,err:=db.Query("select password from user where username=?",username)
	if err!=nil{
		log.Fatal(err)
		return false
	}

	defer stmt.Close()
	for stmt.Next() {
		var row string
		err = stmt.Scan(&row)
		if row==password{
			return true
		}
	}
	return false
}
//判断是否为管理员
func Ai(db * sql.DB) {
	stmt, err := db.Query("select id from userlist where username ='nyx'")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	for stmt.Next() {
		err := stmt.Scan(&Administrator)
		if err != nil {
			log.Fatal(err)
		}

	}
}


//管理员查看积分数
func Select_(db * sql.DB) {
	Ai(db)
	if id==Administrator{
		fmt.Println("欢迎您，管理员")
	}else {
		log.Fatal("对不起，您不是管理员")
	}
	stmt, err := db.Query("select * from signlist;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for stmt.Next() {
		var age int
		var name string

		err := stmt.Scan(&name, &age)
		if err != nil {

			log.Fatal(err)
		}
		fmt.Println(name, age)
	}

}

func Select (c *gin.Context){

}

//管理员增加奖品内容
func Insert (c *gin.Context){

}


//管理员设置奖品积分
func Set (c *gin.Context){

}
