package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func checkErr(err error){
	if err !=nil{
		panic(err)
	}
}

func My_sql(){
	db,err :=sql.Open("mysql","root:password@/test_go?charset=utf8")
	defer db.Close()
	checkErr(err)
	//插入数据
	stmt,err :=db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)
	res,err :=stmt.Exec("张飞","研发部门","2012-12-09")
	checkErr(err)

	id,err := res.LastInsertId()//返回最后插入的记录id
	checkErr(err)
	fmt.Println(id)

	//更新数据
	stmt,err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	res,err = stmt.Exec("吕布",id)
	checkErr(err)

	affect,err := res.RowsAffected()//返回受影响的行
	checkErr(err)
	fmt.Println(affect)

	//查询数据
	rows,err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next(){
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid,&username,&department,&created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt,err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res,err = stmt.Exec(id)
	checkErr(err)
	affect,err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)








}