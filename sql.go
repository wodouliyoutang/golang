package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Tablev struct {
	id     int
	title  string
	route  string
	status int
}

var (
	dbhostsip  = "127.0.0.1" //IP地址
	port       = "3306"      //端口
	dbusername = "root"      //用户名
	dbpassword = "root"      //密码
	dbname     = "laravel"   //库名
	charset    = "utf8"      //字符集
)

var Db *sql.DB

func main() {
	var err error
	Db, err = sql.Open("mysql", dbusername+":"+dbpassword+"@tcp("+dbhostsip+":"+port+")/"+dbname+"?charset="+charset)
	if err != nil {
		log.Fatal(err)
	}
	/*
		crea := Tablev{
			title:  "biaoti",
			route:  "luou",
			status: 2,
		}
		crea.Create()


		del := Tablev{
			id: 4,
		}
		del.Delete()


		sel := Tablev{
			status: 2,
		}
		sel.Select()
	*/
	a, b := RetrievePost(2)
	fmt.Println(a)
	fmt.Println(b)
	defer Db.Close()

}

func (data Tablev) Create() {
	info, err := Db.Exec("INSERT INTO p2p_access (title,route,status) VALUES (?,?,?)", data.title, data.route, data.status)
	if err != nil {
		log.Fatal(err)
	}
	id, err := info.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id)
	return

}

func (data Tablev) Delete() {
	info, err := Db.Exec("DELETE FROM p2p_access WHERE ID = ?", data.id)
	if err != nil {
		log.Fatal(err)
	}
	id, err := info.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id)
	return
}

/**
 * 查
 * @param  {[type]} data Tablev)       Select( [description]
 * @return {[type]}      [description]
 */
func (data Tablev) Select() {
	info, err := Db.Query("SELECT * FROM p2p_access")
	if err != nil {
		log.Fatal(err)
	}
	for info.Next() {
		err := info.Scan(&data.id, &data.title, &data.route, &data.status)
		fmt.Println(err)
		fmt.Println(data.id)
		fmt.Println(data.title)
		fmt.Println(data.route)
		fmt.Println(data.status)
	}
}

/**
 * 查
 * @param {[type]} id int) (post Tablev, err error [description]
 */
func RetrievePost(id int) (post Tablev, err error) {
	post = Tablev{}
	err = Db.QueryRow("SELECT id, title, route FROM p2p_access WHERE id=?", id).Scan(
		&post.id, &post.title, &post.route)
	return
}
