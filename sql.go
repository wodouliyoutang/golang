package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbhostsip  = "127.0.0.1:3306" //IP地址
	dbusername = "root"           //用户名
	dbpassword = "root"           //密码
	dbname     = "laravel"        //库名
	charset    = "utf8"           //字符集
)

func main() {
	db, err := sql.Open("mysql", dbusername+":"+dbpassword+"@tcp("+dbhostsip+")/"+dbname+"?charset="+charset)
	checkErr(err)

	// //插入数据
	// stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	// checkErr(err)

	// res, err := stmt.Exec("码农", "研发部门", "2016-03-06")
	// checkErr(err)

	// id, err := res.LastInsertId()
	// checkErr(err)

	// fmt.Println(id)
	// //更新数据
	// stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	// checkErr(err)

	// res, err = stmt.Exec("码农二代", id)
	// checkErr(err)

	// affect, err := res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM p2p_access")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// //删除数据
	// stmt, err = db.Prepare("delete from p2p_access where uid=?")
	// checkErr(err)

	// res, err = stmt.Exec(id)
	// checkErr(err)

	// affect, err = res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
