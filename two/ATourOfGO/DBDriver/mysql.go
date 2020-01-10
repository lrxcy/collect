package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type TestMysql struct {
	db *sql.DB
}

/* SQL command

CREATE TABLE `atest` (
  `id` bigint(25) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `testnumber` bigint(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

*/

/* 初始化数据库引擎 */
func Init() (*TestMysql, error) {
	test := new(TestMysql)
	db, err := sql.Open("mysql", "root:example@tcp(127.0.0.1:3306)/mysql?charset=utf8")
	if err != nil {
		fmt.Println("database initialize error : ", err.Error())
		return nil, err
	}
	test.db = db
	return test, nil
}

/* 测试数据库数据添加 */
func (test *TestMysql) Create() {
	if test.db == nil {
		return
	}
	stmt, err := test.db.Prepare("insert into atest(id,name,testnumber)values(?,?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	if result, err := stmt.Exec(1, "jim", 123); err == nil {
		if id, err := result.LastInsertId(); err == nil {
			fmt.Printf("insert id : %v\n", id)
		}
	}
	if result, err := stmt.Exec(2, "Jim", 456); err == nil {
		if id, err := result.LastInsertId(); err == nil {
			fmt.Printf("insert id : %v\n", id)
		}
	}
	if result, err := stmt.Exec(3, "JIM", 789); err == nil {
		if id, err := result.LastInsertId(); err == nil {
			fmt.Printf("insert id : %v\n", id)
		}
	}
}

func (test *TestMysql) Close() {
	if test.db != nil {
		test.db.Close()
	}
}

func main() {
	if test, err := Init(); err == nil {
		test.Create()
		test.Close()
	}
}
