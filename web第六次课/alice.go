package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
func main(){
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")

	if err!=nil{
		log.Fatal(err)
	}
	//删除前缀实现函数
	//insertDB(db)
	//deleteDB(db)
	//updateDB(db)
	//selectDB(db)

	}
func insertDB(db *sql.DB)  {
	stmt, err := db.Prepare("insert into nyx(name,age,id) values (?,?,?)")

	if err != nil{
		log.Fatal(err)
	}
	stmt.Exec("Alice","18","1")

}
func deleteDB(db *sql.DB)  {
	stmt, err := db.Prepare("delete from nyx where name = 'Alice'")
	if err != nil{
		log.Fatal(err);
	}
	stmt.Exec();
}
func updateDB(db *sql.DB)  {
	stmt, err := db.Prepare("UPDATE nyx SET name = 'NYX' WHERE id = 'Alice'")
	if err != nil{
		log.Fatal(err)
	}
	stmt.Exec();
}
func selectDB(db *sql.DB)  {
	stmt, err := db.Query("select * from nyx;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for stmt.Next(){
		var id int
		var name string
		var age int
		//var name sql.NullString
		err := stmt.Scan(&name,&age,&id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name,age,id)
	}

}